package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/tina-kuo/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case http.Handler:
		// do nothing, test passed
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, but is %T", v))
	}
}
