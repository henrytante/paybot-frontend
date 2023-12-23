package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("rodando")
	r := router.Gerar()
	err := utils.LoadTemplates()
	if err != nil{
		log.Fatal(err)
	}
	http.ListenAndServe(":3000", r)
}
