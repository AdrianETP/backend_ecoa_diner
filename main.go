package main

import (
	"net/http"
    "ecoadiner.com/v1/routes"
)

func main() {
    http.HandleFunc("/" , routes.HomeRoute)
    http.HandleFunc("/v1" , routes.V1Route)
    http.HandleFunc("/v1/estudiantes" , routes.GetEstudiantes)

    http.ListenAndServe(":8080" , nil)
}
