package parser

import (
	"fmt"
	"regexp"
)

func ParseIntent(msg string) {
	re := regexp.MustCompile(`\*\*(.*?)\*\*`)
	matches := re.FindAllStringSubmatch(msg, -1) // returns the slice of matches -> [-1]

	if len(matches) >0 {
		for _ , m := range matches {
			if len(m) > 1 {
				fmt.Println("🧠 Thought:", m[1])
			}
		}
	}
}