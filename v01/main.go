package main

import "fmt"

func main() {
	fmt.Println("------ 10 about to Make channel --------")
	ch := make(chan int)
	fmt.Println("------ 20 Made channel - about to close channel--------")
	close(ch)
	fmt.Println("------ 30 close channel --------")
	fmt.Println(<-ch)
	fmt.Println("------ 40 waiting for channel ( which is closed ) to return something --------")
}
