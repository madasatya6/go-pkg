package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/madasatya6/go-pkg/websocket"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("online.html")
		if err != nil {
			http.Error(w, "Could not open requested file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// cek online users
		websocket.Connect(w, r, websocket.ONLINE)
	})

	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}
