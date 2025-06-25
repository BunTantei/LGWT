package main

import "fmt"

const englishHelloprefix = "Hello, "

func Hello(name string) string {
	return englishHelloprefix + name
}

func main() {
	fmt.Println(Hello("World"))
}
