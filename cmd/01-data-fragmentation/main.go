package main

import (
	"fmt"
	"githuhub.com/juanjoseluisgarcia/GenLayer/cmd/01-data-fragmentation/model"
	"githuhub.com/juanjoseluisgarcia/GenLayer/cmd/01-data-fragmentation/services"
)

func main() {

	fragments := []model.Fragment{
		{Data: "Hello", Hash: SimpleHash("Hello")},
		{Data: "World", Hash: SimpleHash("World")},
		{Data: "!", Hash: SimpleHash("!")},
	}

	fmt.Println(fragments)

	dataService := services.NewDataService(&HashDataHandler{})

	result := dataService.ReconstructData(fragments)

	fmt.Println(result)
	return
}
