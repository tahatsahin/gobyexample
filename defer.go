package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	// it is important to check errors when closing a file, even in a
	// deferred function
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		panic(err)
	}
}
