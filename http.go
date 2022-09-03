package main

import (
	"log"
	"net/http"
)

func TestGETPlayerss(t *testing.T){
	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()
	
		PlayerServer(response, request)
	
		got := response.Body.String()
		want := "10"
	
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	}

	func PlayerServer(w http.ResponseWriter, r *http.Request) {
		player := strings.TrimPrefix(r.URL.Path, "/players/")
	
		fmt.Fprint(w, GetPlayerScore(player))
	}
	
	func GetPlayerScore(name string) string {
		if name == "Pepper" {
			return "20"
		}
	
		if name == "Floyd" {
			return "10"
		}
	
		return ""
	}