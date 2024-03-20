package main
import (
   "fmt"
   "time"
)
func printIt(msg string) {
   for i := 0; i < 6; i++ {
      time.Sleep(300 * time.Millisecond)
      fmt.Println(msg)
   }
}
func main() {
   // Declare the goroutine
   go printIt("world")
 
   printIt("hello")
}
