package src

import(
	//"fmt"
	//"unicode"
	"strings"
	"regexp"
)

func ValidNumber(s string) bool {
	s = strings.Replace(s, " ", "", -1)//remove speace
	var valid = regexp.MustCompile(`[a-dA-Df-zF-Z]`)
	if valid.MatchString(s) {
		return false
	}
	for i:=0; i<len(s); i++ {
		
	}
	return true
}
