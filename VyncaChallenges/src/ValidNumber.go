package src

import(
	"fmt"
	//"strconv"
	"strings"
	//"regexp"
)

func ValidNumber(s string) bool {
	s = strings.Trim(s, " ")
	fmt.Println(s);
	var pflag, eflag, numflag = false, false, false;
	for i:=0; i<len(s); i++ {
		if('0' <= s[i]&& s[i] <= '9') {
			numflag = true
		} else if(s[i] == '.') {
			if(eflag || pflag) {
				return false
			}
			pflag = true
		} else if(s[i] == 'e') {
			if(eflag || !numflag) {
				return false
			}
			numflag = false
			eflag = true
		} else if(s[i] == '-' || s[i] == '+') {
			if(i != 0 && s[i-1] != 'e'){
				return false
			}
		} else {
			return false
		}
	}
	return numflag;
}
