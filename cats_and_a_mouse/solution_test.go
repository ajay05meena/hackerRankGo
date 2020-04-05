package cats_and_a_mouse_test

import (
	"hackerRankGo/cats_and_a_mouse"
	"testing"
)

func Test_catAndMouse(t *testing.T){
	result := cats_and_a_mouse.CatAndMouse(1, 6, 3)
	if result != "Cat A" {
		t.Error("Wrong answer ", result)
	}

	result = cats_and_a_mouse.CatAndMouse(1, 6, 5)
	if result != "Cat B" {
		t.Error("Wrong answer ", result)
	}

	result = cats_and_a_mouse.CatAndMouse(1, 5, 3)
	if result != "Mouse C" {
		t.Error("Wrong answer ", result)
	}
}