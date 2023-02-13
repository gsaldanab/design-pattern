package main

import (
	"fmt"

	"github.com/gsaldanab/design-pattern/2.factory2/config"
	"github.com/gsaldanab/design-pattern/2.factory2/repository"
)

// factory pattern

func main() {
	config := &config.Configuration{
		Engine: "mysql",
	}

	repo, err := repository.New(config)
	if err != nil {
		panic(err)
	}

	err = repo.Save("hello world")
	if err != nil {
		panic(err)
	}

	data := repo.Find(1)
	fmt.Println(data)
}
