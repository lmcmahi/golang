package main

import (
    "flag"
    "fmt"
)

const APP_VERSION = "0.1"

type  Hello interface {
	hello() int
}

type Person struct {
	age int
	name string
}

//Implements the Hello interface..
func (p *Person) Hello() int {
	fmt.Println(p.name);
	return p.age
}
// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func call(p Person) {
	p.Hello();
}
func main() {
    flag.Parse() // Scan the arguments list
	madi := Person{age:20,name:"coco"};
	pmadi := new(Person);
	pmadi.age=2;
	pmadi.name="co";
	call(madi);
	call(pmadi);

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
}

