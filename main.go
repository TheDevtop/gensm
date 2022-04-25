package main

/*
	Prog: Generate and run state machine
	Vers: 1.0
	Auth: Thijs Haker
*/

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
)

// Loads file, transform to data tape
func loadData(path string) (Tape, error) {
	var err error
	var buf []byte

	if buf, err = os.ReadFile(path); err != nil {
		return nil, err
	}
	data := strings.Fields(string(buf))
	return data, nil
}

// Loads file, transform to text dictionary
func loadText(path string) (Dict, error) {
	var err error
	var buf []byte

	text := make(Dict)

	if buf, err = os.ReadFile(path); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf, &text); err != nil {
		return nil, err
	}
	return text, nil
}

// Runs programmable state machine
func runProgram(text Dict, data Tape, cState string) {
	var (
		nState   string
		routable bool
		stateMap = make(Table)
	)

	for _, in := range data {
		if cState != STATE_HALT {
			if stateMap, routable = text[cState]; !routable {
				log.Fatalf(FMT_SHORT, cState, STATE_NULL)
			}
			if nState, routable = stateMap[in]; !routable {
				log.Fatalf(FMT_BASIC, cState, in, STATE_NULL)
			}
			log.Printf(FMT_BASIC, cState, in, nState)
			cState = nState
			continue
		}
		break
	}
}

func main() {
	var (
		err  error
		text Dict
		data Tape
	)

	textPath := flag.String("t", "program.text", "Specify program text")
	dataPath := flag.String("d", "program.data", "Specify program data")
	mainState := flag.String("m", "Main", "Specify main state")

	flag.Usage = func() {
		println("gensm: Generate and run state machine")
		flag.PrintDefaults()
	}
	flag.Parse()

	if text, err = loadText(*textPath); err != nil {
		log.Panicln(err)
	}
	if data, err = loadData(*dataPath); err != nil {
		log.Panicln(err)
	}

	runProgram(text, data, *mainState)
	os.Exit(0)
}
