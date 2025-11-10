package parser

import (
    "log"
    "regexp"
    "strings"
)

func ParseThoughts(thoughts <-chan string) {
    var partial strings.Builder
    re := regexp.MustCompile(`\*\*(.*?)\*\*`)
    for m := range thoughts {
        partial.WriteString(m)
        fullText := partial.String()
        matches := re.FindAllStringSubmatch(fullText, -1)
        if len(matches) > 0 {
            lastPos := 0
            for _, match := range matches {
                if len(match) > 1 {
                    log.Printf("Thought: %s\n", match[1])
                }
                start := strings.Index(fullText[lastPos:], match[0]) + lastPos
                lastPos = start + len(match[0])
            }
            partial.Reset()
            if lastPos < len(fullText) {
                partial.WriteString(fullText[lastPos:])
            }
        }
    }
    if partial.Len() > 0 {
        matches := re.FindAllStringSubmatch(partial.String(), -1)
        for _, match := range matches {
            if len(match) > 1 {
                log.Printf("Thought: %s\n", match[1])
            }
        }
    }
}