package main

import (

	"github.com/joaofilippe/americanas-desafio/api"
)

func main() {
	api.Server().Run(":8080")
}