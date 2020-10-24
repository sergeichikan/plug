package main

import (
	"fmt"
	"plug/UnitTypes"
	"plugin"
	"time"
)

const mod = "./p/main.so"

func main() {
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("loaded", mod)
	symRun, err := plug.Lookup("CycleRun")
	if err != nil {
		fmt.Println(err)
		return
	}
	run, ok := symRun.(func(<-chan UnitTypes.ChanType, chan<- UnitTypes.ChanType))
	if !ok {
		fmt.Println(ok)
		return
	}
	fmt.Println("loaded CycleRun")

	inputChan := make(chan UnitTypes.ChanType, 1)
	outputChan := make(chan UnitTypes.ChanType, 1)

	go run(inputChan, outputChan)
	go handler(inputChan, outputChan)
	time.Sleep(time.Hour)
}

func handler(inputChan chan<- UnitTypes.ChanType, outputChan <-chan UnitTypes.ChanType) {
	fmt.Println("handler")
	for {
		//t0 := time.Now()
		data := <-outputChan
		for i := 0; i < len(data); i++ {
			data[i].Body = data[i].Body + "change"
		}
		inputChan <- data
		//t1 := time.Now()
		//fmt.Println(t1.Sub(t0))
	}
}
