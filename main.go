package main

import "fmt"

func main() {
    var a chan int
    a = make(chan int)
    fmt.Println(a)
    fmt.Println("Hello Sonar")
    // ToDo: Sonar Bug Test
    b = 5
    if b==0 {
        fmt.Println("bug")
    } else {
	fmt.Println("bug")
    }
}

