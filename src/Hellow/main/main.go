package main

import (
	"awesomeProject"
	"fmt"
	"sync"
	"time"
	"crypto/sha256"
)

type PublicKey interface{}
var (
	route chan int
	wg    sync.WaitGroup
	pWg sync.WaitGroup
	engin machine
	//ping chan <- int // channel to send....
	//pong <-chan int  // channel to receive..
)

func ping(i chan <-int){

}

func pong (o <-chan int){

}

type machine interface {
	velocity() int
	seats() int
}

type car struct {
	speed  int
	nseats int
}

// implements the methods related to machine...

func (m *car) velocity() int {
	return m.speed
}

func (m *car) seats() int {
	return m.nseats
}

// declare a function with poly...

func printMachine(m machine) {
	fmt.Println("speed: ", m.velocity())
}

func changeStruct(c **car) {
	*c = &car{speed: 777, nseats: 11}
	fmt.Printf("@c: %p\n", &*c)
}

func goToRoute(rd chan int) {
	fmt.Println("Going to route..")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Leaving to route..")

	rd <- 1
	defer wg.Done()
}

func exitRoute(rd chan int) int {
	coind := <-rd
	defer wg.Done()
	fmt.Println("Exiting the route")
	return coind
}
func main() {

	fmt.Println("Hellow..")
	mycar := &car{speed: 23, nseats: 23}
	fmt.Printf("mycar: %p", &mycar)
	fmt.Println(mycar)
	changeStruct(&mycar)
	fmt.Println(mycar)

	var newcar *car
	newcar = new(car)
	newcar.nseats = 2222
	newcar.speed = 22
	fmt.Printf("%v\n", &newcar)
	madi := awesomeProject.Person{Age: 23, Name: "coco"}
	fmt.Println(madi)
	route = make(chan int)
	wg.Add(2)
	var coin int
	go goToRoute(route)
	go func() {
		coin = exitRoute(route)
	}()
	wg.Wait()
	fmt.Println("returning coin..",coin)

	fmt.Printf("%T",engin,"...%v",engin)
	engin = new(car)
	fmt.Printf("\n%T",engin,"...%v",engin)

	var hash = sha256.New()
	hash.Write([]byte("Hello world"))
	fmt.Println("\n%x",hash.Sum(nil))

}
