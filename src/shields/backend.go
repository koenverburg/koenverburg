package shields

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type List struct {
	data []Item `json:"data"`
}

type Item struct {
	IconColor  string `json: "iconColor"`
	Label      string `json: "label"`
	LabelColor string `json: "labelColor"`
	Logo       string `json: "logo"`
	LogoColor  string `json: "logoColor"`
	Style      string `json: "style"`
}

// https://img.shields.io/badge/-{iconColor}?style={style}&labelColor={labelColor}&logo={logo}&logoColor={logoColor}&label={label}

func createShieldURL2(s Item) string {
	formattedString := "https://img.shields.io/badge/-%s?style=%s&labelColor=%s&logo=%s&logoColor=%s&label=%s"
	result := fmt.Sprintf(formattedString, s.IconColor, s.Style, s.LabelColor, s.Logo, s.LogoColor, s.Label)
	markdownImage := fmt.Sprintf("![%s Badge](%s)]", s.Label, result)

	return markdownImage
}

func PrepareBackendShields(file string) []string {
	mkImage := make([]string, 1)

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened" + file)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var list List
	json.Unmarshal(byteValue, &list)

	for i := 0; i < len(list.data); i++ {
		shield := createShieldURL2(list.data[i])
		fmt.Sprint(shield)
		mkImage = append(mkImage, shield)
	}

	return mkImage
}
