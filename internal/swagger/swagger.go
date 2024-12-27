package swagger

import (
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"lavka/internal/config"
)

func Setup(cfg config.Swagger, mux *http.ServeMux) error {

	// TODO: кто на ком стоял? а проще можно?

	doc, err := os.ReadFile("docs/tz/openapi.json")
	if err != nil {
		return err
	}

	mux.Handle("GET /swagger/doc.json", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.Write(doc) }))
	mux.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))
	return nil
}
