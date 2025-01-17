package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"lavka/internal/logger"
)

type Validator interface {
	Validate() error
}

type Helper struct {
	w  http.ResponseWriter
	r  *http.Request
	gr string
	op string
	l  *slog.Logger
}

func New(w http.ResponseWriter, r *http.Request, gr, op string) *Helper {
	return &Helper{
		w:  w,
		r:  r,
		gr: gr,
		op: op,
	}
}

func (x *Helper) Ctx() context.Context {
	return x.r.Context()
}

func (x *Helper) Log() *slog.Logger {
	if x.l == nil {
		x.l = logger.GetLoggerFromContextOrDefault(x.r.Context()).
			WithGroup(x.gr).With("op", x.op)
	}
	return x.l
}

func (x *Helper) WriteError(err error) {
	switch err := err.(type) {
	case *HTTPError:
		x.w.WriteHeader(err.StatusCode)
		x.WriteResponse(err)
	default:
		x.Log().Error("unhandled error detected", "error", err)
		x.WriteResponse(ErrInternalError)
	}
}

func (x *Helper) WriteResponse(resp any) {
	x.w.Header().Add("content-type", "application/json")

	if err := json.NewEncoder(x.w).Encode(resp); err != nil {
		x.Log().Error("can't write response", "error", err)
		return
	}
}

func (x *Helper) GetID() (uint64, error) {

	s := x.r.PathValue("id")
	if s == "" {

		// this is a logical error, the method should never be called for
		// requests that do not contain ID in the path

		x.Log().Error("id not found in the path", "path", x.r.URL.Path)
		return 0, ErrInternalError
	}

	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		x.Log().Debug("can't parse id", "id", s, "path", x.r.URL.Path)
		return 0, &HTTPError{
			StatusCode: 400,
			Message:    "id must be unsigned integer",
		}
	}

	return v, nil
}

func (x *Helper) ParseBody(req any) error {

	contentType := x.r.Header.Get("content-type")
	if contentType != "application/json" {
		x.Log().Debug("content-type is not application/json", "content-type", contentType)
		return &HTTPError{
			StatusCode: 400,
			Message:    "content-type must be application/json",
		}
	}

	body, err := io.ReadAll(x.r.Body)
	if err != nil {
		x.Log().Error("can't read request body", "error", err)
		return ErrInternalError
	}

	if err := json.Unmarshal(body, req); err != nil {
		x.Log().Debug("can't parse request body", "error", err, "body", body)
		return &HTTPError{
			StatusCode: 400,
			Message:    "request body does not match the schema",
		}
	}

	switch req := req.(type) {
	case Validator:
		if err := req.Validate(); err != nil {
			x.Log().Debug("can't validate request", "error", err, "req", req)
			return &HTTPError{
				StatusCode: 400,
				Message:    fmt.Sprintf("request contains an incorrect value: %v", err),
			}
		}
	default:

		// never allow incoming data without validation

		x.Log().Error(fmt.Sprintf("%T does not implement Validator", req))
		return ErrInternalError
	}

	return nil
}

func (x *Helper) GetLimitOffset(defaultLimit, defaultOffset int) (int, int, error) {
	q := x.r.URL.Query()

	if q.Has("limit") {
		s := q.Get("limit")
		v, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			x.Log().Debug("can't parse limit", "error", err, "limit", s)
			return 0, 0, &HTTPError{
				StatusCode: 400,
				Message:    "limit must be unsigned integer",
			}
		}
		if v == 0 {
			x.Log().Debug("limit is 0", "limit", s)
			return 0, 0, &HTTPError{
				StatusCode: 400,
				Message:    "limit must be > 0",
			}
		}
		defaultLimit = int(v)
	}

	if q.Has("offset") {
		s := q.Get("offset")
		v, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			x.Log().Debug("can't parse offset", "error", err, "offset", s)
			return 0, 0, &HTTPError{
				StatusCode: 400,
				Message:    "offset must be unsigned integer",
			}
		}
		defaultOffset = int(v)
	}

	return defaultLimit, defaultOffset, nil
}
