package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		lang = strings.Replace(lang, `"`, `\"`, -1)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, fmt.Sprintf("<html><body><h1>http://localhost:2024/?lang=en</h1></body><script>lang=\"%s\";</script></html>", lang))
	})

	http.ListenAndServe(":2024", nil)
}
