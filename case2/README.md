# Case 2

This case is intended to be debugged via test.

To start a test debug session, `dlv test ./case2`
To filter the execution to a single test case, you can use `dlv test ./case2 -- -test.run TestDeleteVowels1`

For more information on the arguments passabble to the test binary, see `go help testflag`

# Description

The `TestDeleteVowels` function should return a string with all the vowels removed. There are two test cases in ./`case2_test.go` that should both pass, but one does not.

# Useful commands

`b case2/case2_test.go:8`: set a Break point at line 8 of `case2_test.go`
`s`: Step into a function (or go to next line if the current line is not a function call)
`so`: Step Out of a function
`r`: Restart the current execution, maintaining all the break points that have been set
`p string(buf)`: print the contents of `buf` cast as a string
