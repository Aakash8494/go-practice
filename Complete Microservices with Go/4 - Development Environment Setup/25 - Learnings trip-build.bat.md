trip-build.bat (The Compiler)
This Windows batch script translates raw Go code into a runnable program for a Linux Docker container.

Line-by-line Breakdown:

• `set CGO_ENABLED=0`: Disables external C code dependencies to create a 100% pure Go program. This prevents crashes when dropping the app into lightweight Linux containers (like Alpine).

• `set GOOS=linux`: Tells the Go compiler to build a program for Linux, even if you are running the script on a Windows machine.

• `set GOARCH=amd64`: Compiles for standard 64-bit processors, standard for most cloud servers and Docker containers.

• `go build -o build/trip-service ./services/trip-service/cmd/main.go`: Compiles the main Go file and outputs the final runnable application named `trip-service` into the `build/` folder.

Why is there no file extension? Unlike Windows, which relies on `.exe` to identify programs, Linux relies on a hidden internal security setting called the executable bit. If that switch is on, Linux knows it's a program and runs it, so the Go compiler outputs a raw Linux binary with no extra letters attached.