package src
import (
	//"fmt"
	"testing"
)

type sample struct{
	number string
	result bool
}

func TestValidNumber(t *testing.T) {
	var list [] sample
	list = []sample{sample{"0",true},
	sample{" 0.1 ", true},
	sample{"abc", false},
	sample{"1 a", false},
	sample{"2e10", true},
	sample{". 1", false},
	sample{".1", true},
	sample{"1.", true},
	sample{"e", false},
	sample{"e2", false},
	sample{"3e", false},
	sample{".", false},
	sample{" ", false},
	sample{"..2", false},
	sample{".e2", false},
	sample{"te1", false},
	sample{". 0e1", false},
	sample{"92e1740e91", false},
	sample{"078332e437", true},}

	for _, sample := range list{
		//fmt.Println(sample.number)
		result:=ValidNumber(sample.number)
		if result != sample.result{
			t.Error("sample:", sample.number," retrun:", result," answer is ", sample.result);
		}
	}
}