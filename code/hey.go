package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",
		func (res http.ResponseWriter,
			req *http.Request) {
			fmt.Fprintln( res, `<html>
  <head><title>Página</title>
   <script>console.log('Cargada')</script>
  </head>
  <body><h1>Cargada</h1></body>
</html>`)
		})
	err := http.ListenAndServe(":50403", nil)
}

