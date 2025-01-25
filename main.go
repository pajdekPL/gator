package main

import (
	"fmt"
	"log"

	"github.com/pajdekpl/gator/internal/config"
)

func main() {
	fmt.Println("Hello world")
	config, err := config.Read()
	if err != nil {
		log.Fatalf("error during reading config %v", err)
	}
	fmt.Printf("Read config: %+v\n", config)
	config.SetUser("pajdek")
	fmt.Println(config.DBURL)
	fmt.Println(config.CurrentUserName)

}
