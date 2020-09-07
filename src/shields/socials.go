package shields

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

func createShieldURL(s Social) string {
	formattedString := "https://img.shields.io/badge/-%s-%s?style=%s&labelColor=%s&logo=%s&logoColor=%s&link=%s"
	result := fmt.Sprintf(formattedString, s.Handle, s.IconColor, s.Style, s.LabelColor, s.Logo, s.LogoColor, s.Link)
	markdownImage := fmt.Sprintf("[![%s Badge](%s)](%s)", s.Platform, result, s.Link)

	return markdownImage
}

func PrepareSocialShields(file string) []string {
	mkImage := make([]string, 5)

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened" + file)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var socials SocialNodes
	json.Unmarshal(byteValue, &socials)

	for i := 0; i < len(socials.Socials); i++ {
		shield := createShieldURL(socials.Socials[i])
		mkImage = append(mkImage, shield)
	}

	return mkImage
}
