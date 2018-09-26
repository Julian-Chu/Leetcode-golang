package perfectNumber

import "testing"

func TestPerfectNumber_1(t *testing.T) {
	res := checkPerfectNumber(28)

	if res != true {
		t.Error("Failed")
	}
}
