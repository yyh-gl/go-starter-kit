package main

import (
	"fmt"
	nethttp "net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yyh-gl/go-starter-kit/app/presetation/http"
)

func main() {
	r := httprouter.New()
	r.GlobalOPTIONS = http.WrapHandler(http.PreflightHandler)

	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:8080")
	fmt.Println("========================")
	nethttp.ListenAndServe(":3000", r)
}
