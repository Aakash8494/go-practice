# Go Testing & Modules: Complete Learning Summary

This document summarizes all the core concepts learned for setting up, writing, and running tests in Go, along with how Go Modules work.

## 1. Project Setup: Go Modules
Before running any tests, Go needs to know the folder is an official project.
* **`go mod init <module-name>`**: The Go equivalent of `npm init`. It creates a `go.mod` file (similar to `package.json`). 
  * *Naming:* It is best practice to use a GitHub URL (e.g., `github.com/username/project`) to prevent naming collisions globally, but local names like `go-practice` work fine for learning.
  * **Note:** This is completely different from `git init` (which tracks code history). `go mod init` handles package management.
* **`go mod tidy`**: The Go equivalent of `npm install`. It scans your code, figures out what packages you need, and downloads them.

## 2. File Naming & Function Structure
* **File Naming:** Test files must end with `_test.go` (e.g., `main_test.go`). They generally live in the exact same folder as the code they are testing.
* **Function Naming:** Test functions must start with the word `Test` followed by a capital letter (e.g., `func TestProcessTruck(t *testing.T)`).

## 3. Writing Tests (Manual Assertions)
Go does not have built-in assertion libraries like JavaScript's `expect(val).toBe(0)`.
* **Standard `if` statements:** You manually verify outcomes using basic logic.
* **`t.Fatalf()`:** If an outcome is wrong, you call `t.Fatalf("Custom error message")`. This immediately stops and fails the test while logging your message.

## 4. Running Tests (The Terminal Commands)
Windows PowerShell handles commands differently than Mac/Linux. **Never use `go test -v *.go` in PowerShell**, because it passes the literal `*` character instead of the file names, causing a syntax error.

* **`go test -v .`** * Runs tests **only** in the current directory you are sitting in.
* **`go test -v ./...`**
  * Runs tests in the current directory **AND all subdirectories recursively**.

### 💡 Crucial Insight from Terminal Testing:
If you are sitting in a parent folder (like `go-practice`) that has **no** test files directly inside it, running `go test -v .` will just output `? go-practice [no test files]`. 
However, running `go test -v ./...` tells Go to dig deeper. It will find the test files hidden inside subfolders (like `advanced_go/`) and successfully execute them!

## 5. The Real World: Mocking with Interfaces
The primary reason we build **Interfaces** (like the `Truck` interface) instead of just using concrete structs is for testing.
* When your app talks to a real Database, you don't want your tests actually reading/writing to it (it's slow and risky).
* By making your function accept a `Database` **interface**, you can pass in a "Fake" or "Mock" database during your tests. It will satisfy Go's compiler and return dummy data, keeping your tests completely safe and isolated.