package utils

import (
	"fmt"
	"strings"
	"time"
)

func printTemplate(format string, params map[string]interface{}) string {
  for key, val := range params {
    format = strings.Replace(format, "{"+key+"}", fmt.Sprintf("%s", val), -1)
  }
  return format
}

func CreateTemplate(socialsShields []string, frontendShields []string, backendShields []string, devopsShields []string, medium string) string {
	date := time.Now().Format(time.UnixDate)
  line := "\n---\n"

	title := "# Hi there 👋"
	intro := "I'm Koen, and I'm a passionated Software Engineer and automation / devOps enthusiast.\n I like to automate the boring stuff and work on my Neovim configuration."

	stack := "## My Stack"

	frontend := fmt.Sprintf("#### Frontend\n%s", strings.Join(frontendShields, ""))
	backend := fmt.Sprintf("#### Backend\n%s", strings.Join(backendShields, ""))
	devops := fmt.Sprintf("#### DevOps\n%s", strings.Join(devopsShields, ""))

	socials := fmt.Sprintf("#### Connect with me on\n%s", strings.Join(socialsShields, ""))
	footer := "Last updated on: " + date

  params := map[string]interface{}{
    "title": title,
    "intro": intro,

    "line": line,

    "stack": stack,
    "frontend": frontend,
    "backend": backend,
    "devops": devops,

    "line1": line,

    "medium": medium,

    "line2": line,

    "socials": socials,

    "line3": line,

    "footer": footer,
  }

  return printTemplate("{title}\n{intro}\n{line}\n{stack}\n{frontend}\n{backend}\n{devops}\n{line1}\n{medium}{line2}\n{socials}\n{line3}\n{footer}", params)
}
