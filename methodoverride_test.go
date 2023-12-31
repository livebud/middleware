package middleware_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/livebud/middleware"
	"github.com/livebud/mux"
	"github.com/matryer/is"
)

func ok() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func TestNoMethod404(t *testing.T) {
	is := is.New(t)
	values := url.Values{}
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(values.Encode()))
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	router := mux.New()
	router.Patch("/", ok())
	router.Use(middleware.MethodOverride())
	router.ServeHTTP(rec, req)
	is.Equal(rec.Code, 404)
}

func TestPatch200(t *testing.T) {
	is := is.New(t)
	values := url.Values{}
	values.Set("_method", http.MethodPatch)
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(values.Encode()))
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	router := mux.New()
	router.Patch("/", ok())
	router.Use(middleware.MethodOverride())
	router.ServeHTTP(rec, req)
	is.Equal(rec.Code, 200)
}

func TestPatchNoBody404(t *testing.T) {
	is := is.New(t)
	req, err := http.NewRequest(http.MethodPost, "/", nil)
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	router := mux.New()
	router.Patch("/", ok())
	router.Use(middleware.MethodOverride())
	router.ServeHTTP(rec, req)
	is.Equal(rec.Code, 404)
}

func TestPatchNoType404(t *testing.T) {
	is := is.New(t)
	values := url.Values{}
	values.Set("_method", http.MethodPatch)
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(values.Encode()))
	is.NoErr(err)
	rec := httptest.NewRecorder()
	router := mux.New()
	router.Patch("/", ok())
	router.Use(middleware.MethodOverride())
	router.ServeHTTP(rec, req)
	is.Equal(rec.Code, 404)
}

func TestPatchInsensitive200(t *testing.T) {
	is := is.New(t)
	values := url.Values{}
	values.Set("_method", "patch")
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(values.Encode()))
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	router := mux.New()
	router.Patch("/", ok())
	router.Use(middleware.MethodOverride())
	router.ServeHTTP(rec, req)
	is.Equal(rec.Code, 200)
}

func TestDelete200(t *testing.T) {
	is := is.New(t)
	values := url.Values{}
	values.Set("_method", "delete")
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(values.Encode()))
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	router := mux.New()
	router.Delete("/", ok())
	router.Use(middleware.MethodOverride())
	router.ServeHTTP(rec, req)
	is.Equal(rec.Code, 200)
}

func TestPut200(t *testing.T) {
	is := is.New(t)
	values := url.Values{}
	values.Set("_method", "put")
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(values.Encode()))
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	router := mux.New()
	router.Put("/", ok())
	router.Use(middleware.MethodOverride())
	router.ServeHTTP(rec, req)
	is.Equal(rec.Code, 200)
}

func TestGet404(t *testing.T) {
	is := is.New(t)
	values := url.Values{}
	values.Set("_method", "get")
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(values.Encode()))
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	router := mux.New()
	router.Get("/", ok())
	router.Use(middleware.MethodOverride())
	router.ServeHTTP(rec, req)
	is.Equal(rec.Code, 404)
}
