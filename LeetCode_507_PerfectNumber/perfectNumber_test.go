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

func TestPerfectNumber_5(t *testing.T) {
	res := checkPerfectNumber(5)

	if res != false {
		t.Error("Failed")
	}
}

func TestPerfectNumber_1(t *testing.T) {
	res := checkPerfectNumber(1)

	if res != false {
		t.Error("Failed")
	}
}

func TestPerfectNumber_2(t *testing.T) {
	res := checkPerfectNumber(2)

	if res != false {
		t.Error("Failed")
	}
}

func TestPerfectNumber_0(t *testing.T) {
	res := checkPerfectNumber(0)

	if res != false {
		t.Error("Failed")
	}
}
