package main

/*
	Prog: Generate and run state machine
	Vers: 0.1
	Auth: Thijs Haker
*/

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
)

var (
	text Dict
	data Tape

	cState string
	nState string
)

func loadData(path string) error {
	var err error
	var buf []byte

	if buf, err = os.ReadFile(path); err != nil {
		return err
	}
	data = strings.Fields(string(buf))
	return nil
}

func loadText(path string) error {
	var err error
	var buf []byte

	text = make(Dict)

	if buf, err = os.ReadFile(path); err != nil {
		return err
	}
	if err = json.Unmarshal(buf, &text); err != nil {
		return err
	}
	return nil
}

func runProgram() {
	var (
		stateMap = make(Table)
		ok       bool
	)

	for it, in := range data {
		if cState != STATE_HALT {
			if stateMap, ok = text[cState]; !ok {
				log.Fatalf(FMT_ERROR, cState, STATE_NULL)
			}
			if nState, ok = stateMap[in]; !ok {
				log.Fatalf(FMT_BASIC, cState, in, STATE_NULL)
			}
			log.Printf(FMT_BASIC, cState, in, nState)
			cState = nState
			continue
		}
		log.Printf("Program Halted: %d cycles\n", it)
		break
	}
}

func usage() {
	println("gensm: Generate and run state machine")
	flag.PrintDefaults()
}

func main() {
	var err error

	textPath := flag.String("t", "program.text", "Specify program text")
	dataPath := flag.String("d", "program.data", "Specify program data")
	mainState := flag.String("m", "Main", "Specify main state")

	flag.Usage = usage
	flag.Parse()

	if err = loadText(*textPath); err != nil {
		log.Panicln(err)
	}
	if err = loadData(*dataPath); err != nil {
		log.Panicln(err)
	}

	cState = *mainState
	runProgram()
	os.Exit(0)
}
