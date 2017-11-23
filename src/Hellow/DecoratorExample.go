package main

import (
	"fmt"
	"time"
)

// Define a common args & data
type Args map[string]string
type Data map[string]string

// Lets define an interface for a common task...In this case a fetcher

type Fetcher interface {
	Fetch(args Args) (data Data, err error)
}

// lets adds more fonctionalities to the basic fetcher...
// decorated using composition..

type Retrier struct {
	fetcher       Fetcher // basic func..
	numberOfRetry int
	waitTime      time.Duration
}

func (r *Retrier) Fetch(args Args) (Data, error) {
	// Now retry x times and wait before every retry

	for i := 0; i < r.numberOfRetry; i++ {
		//Always start with successfull case
		if data, err := r.fetcher.Fetch(args); err == nil {
			fmt.Println("Retrying again %d now with success", i)
			return data, err
		} else if i < r.numberOfRetry {
			fmt.Println("Failed again..%d\n", r.numberOfRetry)
			return Data{}, nil
		}
		fmt.Println("Now going to sleep..")
		time.Sleep(r.waitTime)
	}
	return Data{},nil
}

// define a dummy structure...

type repository struct{}

// Now implements the fetcher interface..

func (repo *repository) Fetch(args Args) (data Data, err error) {
	//
	if len(args) == 0 {
		return Data{}, fmt.Errorf("Please provide arguments...")
	}
	data = Data{"User": "root", "password": "coco"}
	return data, nil
}
func fun1(ptr *repository){
	fmt.Printf("%T",ptr)

}
func main() {
	args:= Args{"filename":"coco"}
	repo := repository{} //empty repository

	//fmt.Println(repo.Fetch(nil))
	retryRepo := Retrier{fetcher: &repo, waitTime: 3000, numberOfRetry: 3}
	data,err := retryRepo.Fetch(args)
	fmt.Println("value..",data,err)

	nrepo := &repository{}

	fun1(nrepo)
	var arrays [2][2]int
	arrays[1][1] = 1
	fmt.Println(arrays)
}
