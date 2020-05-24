package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	instance := GetInstance()
	fmt.Printf("%p\n", instance)
	instance1 := GetInstance()
	fmt.Printf("%p\n", instance1)
	return
}
