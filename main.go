package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()
	fmt.Println(config.Porta)

	fmt.Println(config.StringConexaoBanco)
	fmt.Println("Rodando APi")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
