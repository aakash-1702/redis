package main;
import (
	"fmt"
	"time"
	"sync"
)

func write(id int , writeWorker <- chan int ,mp map[int]int , mu *sync.Mutex , wg *sync.WaitGroup){

	defer wg.Done()

	
	
	for job := range writeWorker{
		
		mu.Lock()
		fmt.Printf("A job reached a writer with id %d",id)
		
		mp[job] = job
		mu.Unlock()
	}
}

func read(id int , readWorker <- chan int , mp map[int]int , mu *sync.Mutex, wg *sync.WaitGroup){
	defer wg.Done()
	for job := range readWorker{
		
		mu.Lock()
		fmt.Printf("A job reached a reader with id %d and value of job is %d",id,job)
		for key ,value := range mp{
			fmt.Println("Key :", key , "Value :" , value)
		}
		mu.Unlock()
	}
}
func main(){
	readWorker := make(chan int)
	writeWorker := make(chan int)
	mp := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(4)

	go write(1,writeWorker,mp,&mu,&wg)
	go write(2,writeWorker,mp,&mu, &wg)
	go read(1,readWorker,mp,&mu, &wg)
	go read(2,readWorker,mp,&mu, &wg)

	for i := 1 ; i <= 20 ; i++{
		if i % 5  != 0{
			writeWorker <- i
		}else {
			readWorker <- i
		}
	}

	close(readWorker)
	close(writeWorker)
	wg.Wait()

	time.Sleep(10 * time.Second)

}