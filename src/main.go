package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type SocialNodes struct {
	Socials []Social `json:"data"`
}

type Social struct {
	Style      string `json: "style"`
	Platform   string `json: "platform"`
	Logo       string `json: "logo"`
	LogoColor  string `json: "logoColor"`
	LabelColor string `json: "labelColor"`
	IconColor  string `json: "iconColor"`
	Handle     string `json: "handle"`
	Link       string `json: "link"`
}

func createShieldURL(s Social) {
	// https://img.shields.io/badge/-@koenverburg_-1ca0f1?style=flat-square&labelColor=1ca0f1&logo=twitter&logoColor=white&link=https://twitter.com/koenverburg_

	formattedString := "https://img.shields.io/badge/-%s-%s?style=%s&labelColor=%s&logo=%s&logoColor=%s&link=%s"
	result := fmt.Sprintf(formattedString, s.Handle, s.IconColor, s.Style, s.LabelColor, s.Logo, s.LogoColor, s.Link)

	markdownImage := fmt.Sprintf("[![%s Badge](%s)](%s)", s.Platform, result, s.Link)

	fmt.Println(markdownImage)
}

func main() {
	// fmt.Println("Started")

	jsonFile, err := os.Open("../data/socials.json")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var socials SocialNodes
	json.Unmarshal(byteValue, &socials)

	for i := 0; i < len(socials.Socials); i++ {
		createShieldURL(socials.Socials[i])
	}

}
