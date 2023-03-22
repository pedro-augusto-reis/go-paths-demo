package main

import "net/http"

func main() {

	http.HandleFunc("/sistema-a", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Alltech Soluções em Tecnologia.</h1>\n " +
			"<h2>Bem vindo ao sistema A.</h2>\n" +
			"<h3>Cluster São Paulo</h3>"))
	})

	http.HandleFunc("/sistema-b", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Alltech Soluções em Tecnologia.</h1>\n " +
			"<h2>Bem vindo ao sistema B.</h2>\n" +
			"<h3>Cluster São Paulo</h3>"))
	})
	http.ListenAndServe(":8080", nil)
}
