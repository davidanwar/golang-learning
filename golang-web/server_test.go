package golangweb

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: ":9090",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
