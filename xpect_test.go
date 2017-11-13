// For now, this file is just a playground to try error messages and examples
// during development. It will eventually be removed.
package main

import (
	"testing"

	. "github.com/marioizquierdo/xpectify/xpec"
)

func Test_FailingStuff(t *testing.T) {
	princess := "Toad"
	Xpect(princess).ToBe("Peach")
}
