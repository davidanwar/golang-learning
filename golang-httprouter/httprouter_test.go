package golanghttprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World!")
	})

	request := httptest.NewRequest("GET", "http://localhost:9090/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello World!", string(body))
}

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		fmt.Fprint(writer, id)
	})

	request := httptest.NewRequest("GET", "http://localhost:9090/products/123", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "123", string(body))
}
