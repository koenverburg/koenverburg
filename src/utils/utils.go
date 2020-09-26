package utils

import (
	"fmt"
	"strings"
	"time"
)

func CreateTemplate(socialsShields []string, frontendShields []string, backendShields []string, devopsShields []string, medium string) string {
	date := time.Now().Format(time.UnixDate)
	title := "# Hi there 👋"
	intro := "I'm Koen Verburg and I'm a Full-Stack Developer and automation and devOps enthusiast."
	socials := fmt.Sprintf("#### Connect with me on\n%s", strings.Join(socialsShields, ""))
	stack := "## My Stack"
	frontend := fmt.Sprintf("#### Frontend \n%s", strings.Join(frontendShields, ""))
	backend := fmt.Sprintf("#### Backend \n%s", strings.Join(backendShields, ""))
	devops := fmt.Sprintf("#### DevOps \n%s\n===\n", strings.Join(devopsShields, ""))
	footer := "Last updated on: " + date

	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n", title, intro, socials, stack, frontend, backend, devops, medium, footer)
}
