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

func TestPanicRouter(t *testing.T) {

	router := httprouter.New()
	router.PanicHandler = func(writer http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(writer, "Panic: ", error)
	}
	router.GET("/products/:id", func(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
		panic("Wah")
	})

	request := httptest.NewRequest("GET", "http://localhost:9090/products/123", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Panic: Wah", string(body))

}
