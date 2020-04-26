# Case 3

To start a debug session, `dlv debug .` or `dlv debug ./case3` from the root.

Alternatively, compile the binary with `go build -o ./case3/run ./case3`, then `dlv exec ./case3/run`.

# Description

This is intended to demonstrate how to work with multiple goroutines and stacks in the debugger.

The `GenerateFibonacci` function will generate the fibonacci series into the the channel it returns. The main function consumes from that channel and panics if the numbers it consumes are not monotonically increasing (spoiler: it will always panic).

# Useful commands

The debugger acts as though it hits a breakpoint if execution panics, so no need to set a breakpoint for this one.

`goroutines`: list the currently running goroutines. Can also shorten to `grs`
`gr 5`: switch to goroutine #5 (numbered in the output of `grs`)
`ls`: print the source code lines surrounding your current location
`stack`: print the current call stack as well as your location in it
`frame 5`: switch to frame #5 in the call stack (frames numbered in the output of `stack`)
