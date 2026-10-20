package main
import ("fmt")
func channel(ch chan string){
	ch <- "Hello World"
}

func main(){
	ch := make(chan string)
	// create a routine which call and perform operation and provide data to channel
	go channel(ch)

	message := <- ch

	fmt.Println(message);
}