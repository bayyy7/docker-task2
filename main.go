package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Someone hit me!")
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Wise Words</h1><p>The greatest wisdom is in simplicity.</p>")
	})

	fmt.Println("Server starting on port 80...")
	http.ListenAndServe(":80", nil)
}
