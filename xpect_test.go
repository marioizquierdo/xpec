package main

import (
	"testing"

	. "code.justin.tv/edge/xpectify/xpect"
)

func Test_Movidas(t *testing.T) {
	v := 1
	Xpect(t, v).ToBe(2)

	s := "foo"
	Xpect(t, s).ToBe("bar")
}
