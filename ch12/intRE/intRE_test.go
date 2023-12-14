package intRE

import (
	"math/rand"
	"strconv"
	"testing"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func Test_matchInt(t *testing.T) {
	if matchInt("") {
		t.Error(`matchInt("") != true`)
	}

	if matchInt("00") == false {
		t.Error(`matchInt("00") != true`)
	}

	if matchInt("-00") == false {
		t.Error(`matchInt("-00") != true`)
	}

	if matchInt("+00") == false {
		t.Error(`matchInt("+00") != true`)
	}
}

func Test_with_random(t *testing.T) {
	n := strconv.Itoa(random(-100000, 19999))

	if matchInt(n) == false {
		t.Error("n = ", n)
	}
}
