package intRE

import (
	"math/rand"
	"strconv"
	"testing"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func TestMatchInt(t *testing.T) {
	if matchInt("") {
		t.Error(`matchInt("") != false`)
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

func TestWithRandom(t *testing.T) {
	n := strconv.Itoa(random(-100000, 19999))

	if matchInt(n) == false {
		t.Error("n = ", n)
	}
}
