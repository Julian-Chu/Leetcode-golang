package perfectNumber

import "testing"

func TestPerfectNumber_28(t *testing.T) {
	res := checkPerfectNumber(28)

	if res != true {
		t.Error("Failed")
	}
}

func TestPerfectNumber_6(t *testing.T) {
	res := checkPerfectNumber(6)

	if res != true {
		t.Error("Failed")
	}
}
