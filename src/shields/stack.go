package shields

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/koenverburg/koenverburg/src/types"
)

func createShieldURL2(s types.Item) string {
	formattedString := "https://img.shields.io/badge/-%s?style=%s&labelColor=%s&logo=%s&logoColor=%s&label=%s"
	result := fmt.Sprintf(formattedString, s.IconColor, s.Style, s.LabelColor, s.Logo, s.LogoColor, s.Label)
	markdownImage := fmt.Sprintf("![%s Badge](%s)", s.Label, result)

	return markdownImage
}

func PrepareStackShields(file string) []string {
	var markdownImage []string

	// abstract this to an util func
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	fmt.Println("Successfully Opened" + file)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var list types.List
	json.Unmarshal(byteValue, &list)

	for i := 0; i < len(list.Data); i++ {
		shield := createShieldURL2(list.Data[i])
		markdownImage = append(markdownImage, shield)
	}

	return markdownImage
}
