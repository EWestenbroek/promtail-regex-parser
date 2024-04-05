package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "<190>"
	pattern := "<(?P<logcode>[0-9]{1,3})>"
	groupNames := []string{
		"logcode",
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Invalid regex pattern")
		return
	}

	match := re.FindStringSubmatch(text)
	if match != nil {
		fmt.Println(match)
		fmt.Println()

		for _, groupName := range groupNames {
			fmt.Printf("%s:\n", groupName)
			fmt.Println(match[re.SubexpIndex(groupName)])
			fmt.Println()
		}
		fmt.Println()
	} else {
		fmt.Println("Wrong regex")
	}
}
