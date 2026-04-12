package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	str := "one"
	fmt.Printf("str: %s\n", str)
	fmt.Printf("&str: %p\n", &str)
	fmt.Printf("data address of str: %p\n", unsafe.StringData(str))

	str = "two"
	fmt.Printf("str: %s\n", str)
	fmt.Printf("&str: %p\n", &str)
	fmt.Printf("data address of str: %p\n", unsafe.StringData(str))

	strHeap := strings.Repeat("A", 5)
	fmt.Println("strHeap before:", strHeap)

	ptrHeap := unsafe.StringData(strHeap)
	dataHeap := unsafe.Slice(ptrHeap, len(strHeap))

	dataHeap[0] = 'B' // This will NOT crash because it's heap memory!

	fmt.Println("strHeap after: ", strHeap)

	// 1. Create a string (allocated on the heap for this example)
	strStack := "Hello"
	fmt.Println("strStack before:", strStack)

	// 2. Get the pointer to the underlying data
	ptrStack := unsafe.StringData(strStack)

	// 3. Create a byte slice that points to the same memory
	// We use unsafe.Slice to treat that pointer as a mutable array
	dataStack := unsafe.Slice(ptrStack, len(strStack))

	// 4. Modify the 'immutable' memory
	dataStack[0] = 'J'

	// 5. Look at the original string variable
	fmt.Println("strStack after: ", strStack)
}
