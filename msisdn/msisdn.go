package main

import (
	"regexp"
	"strings"
)

func normalize(msisdn string) string {
	redigit := regexp.MustCompile(`\D+`)
	n := redigit.ReplaceAllString(msisdn, "")

	if strings.HasPrefix(n, "0") {
		n = strings.Replace(n, "0", "62", 1)
	}
	mayStartWith := []string{"8", "21"}
	for _, token := range mayStartWith {
		if strings.HasPrefix(n, token) {
			n = strings.Replace(n, token, "628", 1)
		}
	}

	return n
}
func main() {

}
