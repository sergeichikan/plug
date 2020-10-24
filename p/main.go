package main

import (
	"fmt"
	"plug/UnitTypes"
	"time"
)

// go build -buildmode=plugin -o p/main.so p/main.go

const cycleLimit = time.Millisecond * 100
const dataLength = 100_000

var Data []UnitTypes.Message // можно передать напрямую
//var InputChan = make(chan []UnitTypes.Message, 1)
//var OutputChan = make(chan []UnitTypes.Message, 1)

//func main() {
//	CycleRun()
//}

func init() {
	for i := 0; i < dataLength; i++ {
		Data = append(Data, UnitTypes.Message{"asdas", "b", time.Now().Unix(), 1.2})
	}
	fmt.Println(len(Data))
}

func iteration(inputChan <-chan UnitTypes.ChanType, outputChan chan<- UnitTypes.ChanType) {
	//time.Sleep(time.Millisecond * 50)
	outputChan <- Data
	Data = <-inputChan
}

func CycleRun(inputChan <-chan UnitTypes.ChanType, outputChan chan<- UnitTypes.ChanType) {
	fmt.Println("CycleRun")
	for {
		fmt.Println("---")
		startTime := time.Now()
		iteration(inputChan, outputChan)
		workTime := time.Since(startTime)
		fmt.Println(workTime)
		sleepTime := cycleLimit - workTime
		fmt.Println(sleepTime)
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		} else {
			fmt.Println("*****************************************************************************************")
		}
	}
}
