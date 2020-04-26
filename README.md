# Getting started with the Delve debugger

The purpose of this is to put together a few examples of working with a `dlv` to get people bootstrapped.

Using the debugger is not hard, but usually by the time we feel we need a debugger, it's because we're desperately trying to solve an issue and we don't want to spend the time climbing the learning curve.

# Why use a debugger?

Personally, I still love `print` and `fmt.Printf` statements. They're super quick to drop in exactly where I'm already looking at the code, and they're pretty flexible. Sometimes, they just don't quite cut it though.

1) Debugging into code you don't control. Sometimes we need to trace execution into a library, and editing the source code with a `print` isn't really an option. Debuggers are great here.
2) Debugging a compiled and/or _running_ binary. Sometimes, the issues only show up on a certain platform, or only show up occasionally in production. In this case, we can't just recompile and print a few lines, so we can attach a debugger to the pre-compiled binary.
3) Stdout is already consumed by something else (e.g. a logger)
4) You really need a detailed step-through of execution. Debuggers are very good at this, and are super helpful for skipping over sections you're not interested in.

Also, note the `edit` command for if you really want to drop back to adding a `print`. It will open up your editor at the current debug location. Super useful for things like taking advantage of code navigation, etc.
