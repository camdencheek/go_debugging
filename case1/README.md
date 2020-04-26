# Case 1

To start a debug session, `dlv debug .` or `dlv debug ./case1` from the root.

Alternatively, compile the binary with `go build -o ./case1/run ./case1`, then `dlv exec ./case1/run`.

## Description

This program should allocate `a` with capacity 5, then fill that slice with the elements of `b`.
It should print `[]int{0, 1, 2, 3, 4}`.

## Useful commands

`h`: print Help, which gives you a list of commands
`b main.main`: set a Break point at the beginning of the main function
`c`: Continue to the next break point
`n`: advance to the Next line
`p a`: Print a debug representation of the variable `a`
