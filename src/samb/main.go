package main

import (
	"fmt"
	"log"
	"net/http"
	"samb/routes"
)

func main()  {
	r := routes.NewRouter()
	fmt.Println("Aplication Running in port :8080")
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
