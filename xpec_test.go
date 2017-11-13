// For now, this file is just a playground to try error messages and examples
// during development. It will eventually be removed.
package main

import (
	"testing"

	. "github.com/marioizquierdo/xpec/e"
)

func Test_FailingStuff(t *testing.T) {
	princess := "Toad"
	Xpec(t, princess).ToBe("Peach")
}
