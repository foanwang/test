package src
import (
	"testing"
)

type sample struct{
	number string
	result bool
}

func TestValidNumber(t *testing.T) {
	var list [] sample
	list = []sample{sample{"0",true}, sample{" 0.1 ", true},sample{"abc", false},sample{"1 a", false},sample{"2e10", true}}

	for _, sample := range list{
		result:=ValidNumber(sample.number)
		if result != sample.result{
			t.Error("sample:", sample.number,"test is worng");
		}
	}
}