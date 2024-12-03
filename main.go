package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Mondoo Engineer")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server läuft auf http://localhost:8080")
	// nosemgrep go.lang.security.audit.net.use-tls.use-tls
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Fehler beim Starten des Servers:", err)
	}
}
