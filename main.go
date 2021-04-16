package main

import (
	"net/http"

	"github.com/pluralsight/webservices/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
