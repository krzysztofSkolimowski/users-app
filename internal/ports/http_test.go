package ports

import (
	"net/http"
	"testing"
)

func TestHttpServer_PostUsers(t *testing.T) {
	tests := []struct {
		name string
		w    http.ResponseWriter
		r    *http.Request
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := HttpServer{}

			r.PostUsers(tt.w, tt.r)
		})
	}
}
