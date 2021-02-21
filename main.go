package main

import (
	"log"
	"net/http"

	"pokemons/routers"
)

func main() {
	router := routers.NewRouters()

	log.Fatal(http.ListenAndServe(":3030", router))
}
