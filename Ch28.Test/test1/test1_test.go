package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) should be 81 but square(9) returns %d", rst)
	}
}

func TestSquare2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("square(3) should be 9 but square(9) returns %d", rst)
	}
}

func TestSquare3(t *testing.T) {
	assert.Equal(t, 81, square(9), "square(9) should be 81")
}

func TestSquare4(t *testing.T) {
	assert.Equal(t, 9, square(3), "square(3) should be 9")
}
