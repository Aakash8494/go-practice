//go:build ignore

package main

import (
	"fmt"
)

// User defines a simple struct to demonstrate general formatting
type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("=== Go 'fmt' Package: Formatting Verbs Cheat Sheet ===")

	// ---------------------------------------------------------
	// 1. General Formatting (The "v" family & types)
	// ---------------------------------------------------------
	fmt.Println("\n--- 1. General Formatting ---")
	u := User{Name: "Alice", Age: 30}

	// %v is the ultimate fallback. It prints the default representation.
	fmt.Printf("%%v  (Default):         %v\n", u)

	// %+v is amazing for structs. It adds the field names!
	fmt.Printf("%%+v (Fields added):    %+v\n", u)

	// %#v prints the value in Go syntax (valid Go code). Great for debugging.
	fmt.Printf("%%#v (Go syntax):       %#v\n", u)

	// %T prints the TYPE of the variable, not the value.
	fmt.Printf("%%T  (Type):            %T\n", u)

	// ---------------------------------------------------------
	// 2. Integers
	// ---------------------------------------------------------
	fmt.Println("\n--- 2. Integers ---")
	num := 42

	fmt.Printf("%%d (Base 10):          %d\n", num) // Standard integer
	fmt.Printf("%%b (Binary):           %b\n", num) // 101010
	fmt.Printf("%%c (Character):        %c\n", 65)  // Output: A (Unicode char)
	fmt.Printf("%%x (Hex lowercase):    %x\n", num) // 2a
	fmt.Printf("%%X (Hex uppercase):    %X\n", num) // 2A

	// ---------------------------------------------------------
	// 3. Floats
	// ---------------------------------------------------------
	fmt.Println("\n--- 3. Floats ---")
	pi := 3.14159265

	fmt.Printf("%%f (Default float):    %f\n", pi)   // 3.141593 (6 decimal places by default)
	fmt.Printf("%%.2f (2 decimal pts):  %.2f\n", pi) // 3.14 (Rounds the number)
	fmt.Printf("%%e (Scientific):       %e\n", pi)   // 3.141593e+00

	// ---------------------------------------------------------
	// 4. Strings & Bytes
	// ---------------------------------------------------------
	fmt.Println("\n--- 4. Strings & Bytes ---")
	text := "Gopher"
	bytes := []byte("Go")

	fmt.Printf("%%s (Plain string):     %s\n", text)  // Gopher
	fmt.Printf("%%q (Quoted string):    %q\n", text)  // "Gopher" (Adds quotes safely)
	fmt.Printf("%%x (Base 16 bytes):    %x\n", bytes) // 476f (Hex dump of bytes)

	// ---------------------------------------------------------
	// 5. Pointers & Booleans
	// ---------------------------------------------------------
	fmt.Println("\n--- 5. Pointers & Booleans ---")
	truth := true
	ptr := &num

	fmt.Printf("%%t (Boolean):          %t\n", truth) // true or false
	fmt.Printf("%%p (Pointer address):  %p\n", ptr)   // 0x... (Prints memory address)
}
