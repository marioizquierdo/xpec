# Xpec

Expectation helpers for Go tests. Similar to testify assertions, but provides error messages that include context information without the need to add assertion messages:

```go
func Test_AnotherCastle(t *testing.T) {
	princess := "Toad"
	Xpec(t, princess).ToBe("Peach")
}
```

running `go test` on this file will output:
```sh
--- FAIL: Test_AnotherCastle (0.00s)
	another_castle_test.go:13
		Xpec(t, princess).ToBe("Peach")
		but "Toad" is not "Peach"
```

The error message includes the code line that is testing the expectation, which shows variable names like `princess`, giving context to the error.

Expectations make clear assertions, helping to produce semantically meaningful errors, which helps developers to quickly identify the "why" of failed tests.
