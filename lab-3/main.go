package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		if strings.Contains(lang,"\\"){
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		lang = Encoding(lang)
		w.Header().Set("Content-Type", "text/html")
		html:=fmt.Sprintf("<html><body><h1>http://localhost:2024/?lang=en</h1></body><script>lang=\"%s\";</script></html>", lang)
		if strings.Contains(html,"%") {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, html)
	})

	http.ListenAndServe(":2024", nil)
}



func Encoding(text string) string {
	text = strings.Replace(text, "\"", "\\\"", -1)
	text = strings.Replace(text, "<", "&lt;", -1)
	text = strings.Replace(text, ">", "&gt;", -1)
	return text
}
