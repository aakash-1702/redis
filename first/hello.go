package main
import (
	"fmt"
	"time"
)

func sayHello(){
	time.Sleep(2 * time.Second)
	fmt.Println("Hello")
}

func main(){
	go sayHello(); // starting a new goroutine is what this go keyword is here for
	fmt.Println("World")
	time.Sleep(3 * time.Second)
}
