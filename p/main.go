package main

import (
	"fmt"
	"plug/UnitTypes"
	"time"
)

// go build -buildmode=plugin -o p/main.so p/main.go

var Data []UnitTypes.Message // можно передать напрямую

func initData(length int) {
	for i := 0; i < length; i++ {
		Data = append(Data, UnitTypes.Message{"name", "body", time.Now().Unix(), 1.2})
	}
	fmt.Println("data length", len(Data))
}

func iteration(inputChan <-chan UnitTypes.ChanType, outputChan chan<- UnitTypes.ChanType) {
	outputChan <- Data
	Data = <-inputChan
}

func Run(
	inputChan <-chan UnitTypes.ChanType,
	outputChan chan<- UnitTypes.ChanType,
	cycleLimit time.Duration,
	dataLength int,
) {
	initData(dataLength)
	for {
		startTime := time.Now()
		iteration(inputChan, outputChan)
		sleepTime := cycleLimit - time.Since(startTime)
		fmt.Println(sleepTime)
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
}
