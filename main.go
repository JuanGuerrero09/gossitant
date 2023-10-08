/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"fmt"
	"mycli/cmd"
	"os"
)






	func check(e error) {
		if e != nil {
			panic(e)
    }
	}
	
	func main() {
	cmd.Execute()

    d1 := []byte("hello\ngo\n")
		fmt.Println(d1)
    err := os.WriteFile("/tmp/dat1", d1, 0644)
    check(err)
		
    f, err := os.Create("/tmp/dat2")
		fmt.Println(f)
    check(err)
		
    defer f.Close()
		
    d2 := []byte{115, 111, 109, 101, 10}
		fmt.Println(d2)
    n2, err := f.Write(d2)
		fmt.Println(n2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

    f.Sync()

    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush()

}