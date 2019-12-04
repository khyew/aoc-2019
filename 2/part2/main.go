package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const OP_HALT = 99
const OP_ADD = 1
const OP_MULT = 2

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && (data[len(data)-1] == '\r' || data[len(data)-1] == '\n') {
		return data[0 : len(data)-1]
	}
	return data
}

// scanCommas is a split function that splits on commas
func scanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func runIntcode(noun, verb int, intcode []int) (int, error) {

	intcode[1] = noun
	intcode[2] = verb

	for pc := 0; pc < len(intcode); pc += 4 {
		switch opcode := intcode[pc]; opcode {
		case OP_HALT:
			//					fmt.Printf("pc(%v) HALT\n", pc)
			//					fmt.Printf("output: %v\n", data[0])
			return intcode[0], nil
		case OP_ADD:
			laddr := intcode[pc+1]
			raddr := intcode[pc+2]
			dest := intcode[pc+3]

			lval := intcode[laddr]
			rval := intcode[raddr]

			//					fmt.Printf("pc(%v) ADD %v(%v), %v(%v) into %v\n", pc, laddr, lval, raddr, rval, dest)
			intcode[dest] = lval + rval
		case OP_MULT:
			laddr := intcode[pc+1]
			raddr := intcode[pc+2]
			dest := intcode[pc+3]

			lval := intcode[laddr]
			rval := intcode[raddr]

			//					fmt.Printf("pc(%v) MULT %v(%v), %v(%v) into %v\n", pc, laddr, lval, raddr, rval, dest)
			intcode[dest] = lval * rval
		}
	}

	return 0, errors.New("ship computer bursts into flames")
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Printf("Couldn't open input file: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(scanCommas)

	program := []int{}

	for scanner.Scan() {
		integer := scanner.Text()
		parsedInteger, err := strconv.Atoi(integer)

		if err != nil {
			fmt.Printf("Couldn't parse input value (%v) into integer: %v\n", integer, err)
		}

		program = append(program, parsedInteger)
	}
	fmt.Printf("Scanned %d comma-separated values into program buffer\n", len(program))

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {

			intcode := make([]int, len(program))
			copy(intcode, program)

			output, err := runIntcode(noun, verb, intcode)
			if err != nil {
				fmt.Printf("computer fault\n")
				return
			}

			if output == 19690720 {
				fmt.Printf("answer: %v\n", 100*noun+verb)
				return
			}
		}
	}
}
