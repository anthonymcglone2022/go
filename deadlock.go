package main
import (
    "fmt"
)
func main() {
    queue := make(chan int, 2)
    queue <- 1
    queue <- 2
    // Program will block at the next line of code
    // If it is removed, however
    // the program will not block
    queue <- 3
 
    fmt.Println(<-queue)
    fmt.Println(<-queue)
}
