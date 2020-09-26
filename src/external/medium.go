package external

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/koenverburg/koenverburg/src/generated"
)

func Medium() string {
	resp, err := http.Get("https://medium.com/feed/@koenverburg")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var rss generated.Rss
	errXML := xml.Unmarshal(body, &rss)
	if errXML != nil {
		log.Fatalln(errXML)
	}

	var items []string
	for i := 0; i < 3; i++ {
		item := fmt.Sprintf("[%s](%s)\n", rss.Channel.Item[i].Title, rss.Channel.Item[i].Link)
		items = append(items, item)
	}

	mediumTitle := fmt.Sprintf("\n### %s\n%s\n", rss.Channel.Title, strings.Join(items, "\n"))
	return mediumTitle
}
