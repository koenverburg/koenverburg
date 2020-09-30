package shields

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/koenverburg/koenverburg/src/types"
)

func createShieldURL(s types.Item) string {
	formattedString := "https://img.shields.io/badge/-%s-%s?style=%s&labelColor=%s&logo=%s&logoColor=%s&link=%s"
	result := fmt.Sprintf(formattedString, s.Handle, s.IconColor, s.Style, s.LabelColor, s.Logo, s.LogoColor, s.Link)
	markdownImage := fmt.Sprintf("[![%s Badge](%s)](%s)", s.Platform, result, s.Link)

	return markdownImage
}

func PrepareSocialShields(file string) []string {
	var markdownImage []string

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened" + file)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var socials types.List
	json.Unmarshal(byteValue, &socials)

	for i := 0; i < len(socials.Data); i++ {
		shield := createShieldURL(socials.Data[i])
		markdownImage = append(markdownImage, shield)
	}

	return markdownImage
}
