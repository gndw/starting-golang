package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gndw/starting-golang/internals/constants"
	"github.com/gndw/starting-golang/internals/services/httpmiddleware"
)

type Implementation struct {
	handler    *http.ServeMux
	middleware httpmiddleware.Service
}

func NewHttpServerService(ctx context.Context, middleware httpmiddleware.Service) (*Implementation, error) {
	handler := http.NewServeMux()
	return &Implementation{handler: handler, middleware: middleware}, nil
}

// https://jsonapi.org/format/#document-structure
type HttpResponse struct {
	Data   interface{}         `json:"data,omitempty"`
	Errors []HttpErrorResponse `json:"errors,omitempty"`
}

type HttpErrorResponse struct {
	Title string `json:"title"`
}

func (m *Implementation) RegisterEndpoint(ctx context.Context, method string, path string, f constants.HttpFunction) error {
	m.handler.HandleFunc(fmt.Sprintf("%v %v", method, path), func(w http.ResponseWriter, r *http.Request) {

		// setup middleware
		hf := m.middleware.LogMiddleware(f)

		// executing function and construct http response
		response, err := hf(r.Context(), w, r)
		httpResponse := HttpResponse{}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			httpResponse.Errors = append(httpResponse.Errors, HttpErrorResponse{Title: err.Error()})
		} else {
			w.WriteHeader(http.StatusOK)
			httpResponse.Data = response
		}

		w.Header().Add("Content-Type", "application/json")
		b, _ := json.Marshal(httpResponse)
		w.Write(b)
	})
	return nil
}

func (m *Implementation) Serve(ctx context.Context) error {
	s := &http.Server{
		Addr:         ":8080",
		Handler:      m.handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("[Http-Service] server starting at port 8080...")
	return s.ListenAndServe()
}
