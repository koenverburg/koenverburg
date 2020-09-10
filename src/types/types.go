package types

type List struct {
	Data []Item `json:"data"`
}

type Item struct {
	Style      string `json: "style"`
	Platform   string `json: "platform"`
	Logo       string `json: "logo"`
	LogoColor  string `json: "logoColor"`
	LabelColor string `json: "labelColor"`
	IconColor  string `json: "iconColor"`
	Handle     string `json: "handle"`
	Link       string `json: "link"`
	Label      string `json: "label"`
}
