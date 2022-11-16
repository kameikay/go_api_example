package main

import (
	"fmt"

	"github.com/kameikay/api_example/configs"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	fmt.Println(configs)
}
