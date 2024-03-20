package main
import "fmt"
func sendToChannel(number int, newChannel chan int) {
    for item := 0; item < number; item++ {
        newChannel <- item
    }
    close(newChannel)
}
func main() {
    // Unbuffered channel
    newChannel := make(chan int)
 
    // Function to send six numbers to the channel
    go sendToChannel(6, newChannel)
    // Receive items from the channel
    for item := range newChannel { 
        fmt.Println(item)
    }
}
