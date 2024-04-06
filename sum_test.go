package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error("O resultado esperado é diferente de 5")
	}
}

func TestSub(t *testing.T) {

	result := sub(10, 6)

	if result != 4 {
		t.Error("O resultado esperado é diferente de 4")
	}
}
