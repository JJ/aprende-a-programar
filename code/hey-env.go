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
				fmt.Fprintln( res, `<html>
  <head><title>Página</title></head>
  <body><h1>Cargada</h1></body>
</html>`)
			})
	ε := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if ε != nil {
		panic( ε )
	}
}

