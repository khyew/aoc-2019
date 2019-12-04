package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const OP_HALT = 99
const OP_ADD = 1
const OP_MULT = 2

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
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

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Printf("Couldn't open input file: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(scanCommas)

	data := []int{}

	for scanner.Scan() {
		integer := scanner.Text()
		parsedInteger, err := strconv.Atoi(integer)

		if err != nil {
			fmt.Printf("Couldn't parse input value (%v) into integer: %v\n", integer, err)
		}

		data = append(data, parsedInteger)
	}
	fmt.Printf("Scanned %d comma-separated values into data buffer\n", len(data))

	for pc := 0; pc < len(data); pc += 4 {
		switch opcode := data[pc]; opcode {
		case OP_HALT:
			fmt.Printf("pc(%v) HALT\n", pc)
			fmt.Printf("output: %v\n", data[0])
			return
		case OP_ADD:
			laddr := data[pc+1]
			raddr := data[pc+2]
			dest := data[pc+3]

			lval := data[laddr]
			rval := data[raddr]

			fmt.Printf("pc(%v) ADD %v(%v), %v(%v) into %v\n", pc, laddr, lval, raddr, rval, dest)
			data[dest] = lval + rval
		case OP_MULT:
			laddr := data[pc+1]
			raddr := data[pc+2]
			dest := data[pc+3]

			lval := data[laddr]
			rval := data[raddr]

			fmt.Printf("pc(%v) MULT %v(%v), %v(%v) into %v\n", pc, laddr, lval, raddr, rval, dest)
			data[dest] = lval * rval
		}
	}
}
