package main

import (
	"io"
	"os"
	"fmt"

	"github.com/koenverburg/koenverburg/src/utils"
	"github.com/koenverburg/koenverburg/src/shields"
)

func main() {
	fmt.Println("Started")

	socialShields := shields.PrepareSocialShields("../data/socials.json")
	devopsShields := shields.PrepareStackShields("../data/stack.devops.json")
	backendShields := shields.PrepareStackShields("../data/stack.backend.json")
	frontendShields := shields.PrepareStackShields("../data/stack.frontend.json")

	template := utils.CreateTemplate(socialShields, frontendShields, backendShields, devopsShields)

	file, err := os.Create("../README.md")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = io.WriteString(file, template)
	if err != nil {
		fmt.Println(err)
	}

	file.Sync()
}
