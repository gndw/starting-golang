package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gndw/starting-golang/internals/constants"
	"github.com/gndw/starting-golang/internals/services/env"
	"github.com/gndw/starting-golang/internals/services/httpmiddleware"
)

type Implementation struct {
	handler    *http.ServeMux
	middleware httpmiddleware.Service
	env        env.Service
	server     *http.Server
}

func NewHttpServerService(ctx context.Context, middleware httpmiddleware.Service, env env.Service) (*Implementation, error) {
	handler := http.NewServeMux()
	return &Implementation{handler: handler, middleware: middleware, env: env}, nil
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
		m.writeResponse(w, response, err)
	})
	return nil
}

func (m *Implementation) writeResponse(w http.ResponseWriter, response interface{}, err error) {
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
}

func (m *Implementation) Start(ctx context.Context) error {
	port := m.env.Get(ctx).Port
	if port == "" {
		return fmt.Errorf("port is empty")
	}
	m.server = &http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		Handler:      m.handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("[Http-Service] server starting at port %v...\n", port)
	go func() {
		if err := m.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("[Http-Service] server failed to start: %v\n", err)
		}
	}()
	return nil
}

func (m *Implementation) Shutdown(ctx context.Context) error {
	if m.server == nil {
		return nil
	}
	fmt.Printf("[Http-Service] server shutting down...\n")
	return m.server.Shutdown(ctx)
}
