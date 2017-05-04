package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/",
		func (res http.ResponseWriter,
			req *http.Request) {
				fmt.Fprintln( res, MainPage())
			})
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "50403"
	}
	ε := http.ListenAndServe(":"+port, nil)
	if ε != nil {
		panic( ε )
	}
}

func MainPage() string {
	return `<html>
  <head><title>Página</title>
   <script>console.log('☒ Cargada')</script>
  </head>
  <body><h1>Cargada</h1></body>
</html>`
}

