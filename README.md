# Xpectify

Expectation helpers for Go tests. Similar to testify assertions, but provides error messages that can be easier to read without the need to include assertion messages:

```go
func Test_AnotherCastle(t *testing.T) {
	princess := "Toad"
	Xpect(t, princess).ToBe("Peach")
}
```

running `go test` on this file will output:
```sh
--- FAIL: Test_AnotherCastle (0.00s)
	comparison_test:11
		Xpect(t, princess).ToBe("Peach")
		expected "Toad" to be "Peach"
```

The error message includes the line that is testing the expectation, which includes variable names like `princess` on it, giving context to the message.

Expectations make it clear what is being compared, which helps producing a semantically meaningful error messages, which helps developers identify the "why" on failed tests faster.