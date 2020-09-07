package main

import (
	"fmt"

	"github.com/koenverburg/readme/shields"
)

func main() {
	fmt.Println("Started")

	socialShields := shields.PrepareSocialShields("../data/socials.json")
	fmt.Println(socialShields)
}
