package main
import (
	"fmt"
	"time"
)

func doingJob(id int , jobs <- chan int){
	for job := range jobs{
	    fmt.Printf("Job of id %d is done by worker %d\n",job,id)
		time.Sleep(time.Second)
	}
}

func main(){
	jobs := make(chan int)
	
    go doingJob(2, jobs)
	go doingJob(1 , jobs)
	

	for job := 1 ; job <= 4 ; job++ {
		jobs <- job
	}

	close(jobs)
	time.Sleep(4 * time.Second)
}