package main

import (
	"flag"
	"fmt"
	"plug/UnitTypes"
	"plugin"
	"time"
)

const defaultPluginPath = "./p/main.so"
const defaultCycleLimit = time.Millisecond * 100
const defaultDataLength = 100_000

func main() {
	pluginPath := flag.String("plugin", defaultPluginPath, "path to plugin")
	cycleLimit := flag.Int("cycle", int(defaultCycleLimit), "time limit of iteration")
	dataLength := flag.Int("length", defaultDataLength, "data length")
	flag.Parse()
	fmt.Println("plugin", *pluginPath)
	fmt.Println("cycle", time.Duration(*cycleLimit))
	fmt.Println("length", *dataLength)
	runPlugin(*pluginPath, time.Duration(*cycleLimit), *dataLength)
}

func runPlugin(pluginPath string, cycleLimit time.Duration, dataLength int) {
	plug, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("loaded", pluginPath)

	symRun, err := plug.Lookup("Run")
	if err != nil {
		fmt.Println(err)
		return
	}

	run, ok := symRun.(func(<-chan UnitTypes.ChanType, chan<- UnitTypes.ChanType, time.Duration, int))
	if !ok {
		fmt.Println(ok)
		return
	}
	fmt.Println("loaded Run")

	inputChan := make(chan UnitTypes.ChanType, 1)
	outputChan := make(chan UnitTypes.ChanType, 1)

	go run(inputChan, outputChan, cycleLimit, dataLength)
	go handler(inputChan, outputChan)

	time.Sleep(time.Hour)
}

func handler(inputChan chan<- UnitTypes.ChanType, outputChan <-chan UnitTypes.ChanType) {
	for {
		data := <-outputChan
		timestamp := time.Now().Unix()
		for i := range data {
			data[i].Time = timestamp
		}
		inputChan <- data
	}
}
