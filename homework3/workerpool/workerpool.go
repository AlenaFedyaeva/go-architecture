package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Results struct{
	CountBadRequest int64
	CountRequest int64
	Mu *sync.RWMutex
	ArrTimes []time.Duration
}

var once sync.Once
var results *Results

func GetResults() *Results {

    once.Do(func() {
        results = &Results{
			Mu: &sync.RWMutex{},
		}
    })
    return results
}

func (rez *Results) getAverageTime() (float64,error){
	rez.Mu.RLock()
	var sum float64 
	for _, val := range rez.ArrTimes {
		sum += float64(val)
	}
	rez.Mu.RUnlock()

	// check 0
	del:=float64(rez.CountRequest-rez.CountBadRequest)
	if del == 0 {
		return 0, errors.New("Dividing by zero")
	}

	avg:=float64(sum)/del
	return avg,nil
}


type Job struct {
	request *http.Request
}

type Worker struct {
	wg      *sync.WaitGroup
	num     int // only for example
	jobChan <-chan *Job
	
}

func (w *Worker) Handle() {
	defer w.wg.Done()
	for job := range w.jobChan {


		request := job.request
		log.Printf("worker %d processing job", w.num, request)

		timeStart := time.Now()
		resp, err := http.DefaultClient.Do(request)
		timeDuration := time.Now().Sub(timeStart)

		results.Mu.Lock()
		results.CountRequest++;
		if err != nil {
			results.CountBadRequest++;
			fmt.Println(err)
		}
		results.ArrTimes=append(results.ArrTimes, timeDuration)
		results.Mu.Unlock()
		if resp!=nil{
			resp.Body.Close()
		}
		
		fmt.Println("duration ", timeDuration)

	}
}

func NewWorker(num int, wg *sync.WaitGroup, jobChan <-chan *Job) *Worker {
	return &Worker{
		wg:      wg,
		num:     num,
		jobChan: jobChan,
	}
}
