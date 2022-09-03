package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)
// func ListenAndServer(addr string, handler Handler) error

// type Handler interface{
// 	ServerHttp(ResponseWriter, *Request)
// }
func PlayerServer(w http.ResponseWriter, r *http.Request){
     fmt.Fprint(w,"20")
}
func TestGETPlayerss(t *testing.T){
	t.Run("returns Papper's score", func(t *testing.T){

		request, _ := http.NewRequest(http.MethodGet, "/players/Papper",nil)

		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want{
			t.Errorf("got %q, want %q",got, want)
		}
	})
}

