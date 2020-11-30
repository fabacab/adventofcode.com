package main

// --- Part Two ---
//
// "Good, the new computer seems to be working correctly! Keep it nearby during this mission - you'll probably use it again. Real Intcode computers support many more features than your new one, but we'll let you know what they are as you need them."
//
// "However, your current priority should be to complete your gravity assist around the Moon. For this mission to succeed, we should settle on some terminology for the parts you've already built."
//
// Intcode programs are given as a list of integers; these values are used as the initial state for the computer's memory. When you run an Intcode program, make sure to start by initializing memory to the program's values. A position in memory is called an address (for example, the first value in memory is at "address 0").
//
// Opcodes (like 1, 2, or 99) mark the beginning of an instruction. The values used immediately after an opcode, if any, are called the instruction's parameters. For example, in the instruction 1,2,3,4, 1 is the opcode; 2, 3, and 4 are the parameters. The instruction 99 contains only an opcode and has no parameters.
//
// The address of the current instruction is called the instruction pointer; it starts at 0. After an instruction finishes, the instruction pointer increases by the number of values in the instruction; until you add more instructions to the computer, this is always 4 (1 opcode + 3 parameters) for the add and multiply instructions. (The halt instruction would increase the instruction pointer by 1, but it halts the program instead.)
//
// "With terminology out of the way, we're ready to proceed. To complete the gravity assist, you need to determine what pair of inputs produces the output 19690720."
//
// The inputs should still be provided to the program by replacing the values at addresses 1 and 2, just like before. In this program, the value placed in address 1 is called the noun, and the value placed in address 2 is called the verb. Each of the two input values will be between 0 and 99, inclusive.
//
// Once the program has halted, its output is available at address 0, also just like before. Each time you try a pair of inputs, make sure you first reset the computer's memory to the values in the program (your puzzle input) - in other words, don't reuse memory from a previous attempt.
//
// Find the input noun and verb that cause the program to produce the output 19690720. What is 100 * noun + verb? (For example, if noun=12 and verb=2, the answer would be 1202.)

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
    "fmt"
)

// readIntcodes loads an intcode program and returns a slice of ints.
func readIntcodes() []int {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	sourceCode, _ := r.Read()

	var prog []int
	for _, v := range sourceCode {
		val, _ := strconv.Atoi(v)
		prog = append(prog, val)
	}

	return prog
}

// run executes the intcode program with the given noun and verb.
func run(noun int, verb int) int {
	const opcodeLength = 4

	prog := readIntcodes()

	// Set the initial program state as per challenge instructions.
	prog[1] = noun
	prog[2] = verb

	// Process program by starting the instruction pointer (`ip`) at 0
	// (the start), then step forward `opcodeLength` per instruction.
evaluation: // Label this `for` loop so we can `break` out of it.
	for ip := 0; ip < len(prog); ip += opcodeLength {
		switch prog[ip] {

		case 1: // opcode `1` means addition.
			addend1pos := ip + 1
			addend2pos := ip + 2
			sumPos := ip + 3
			prog[prog[sumPos]] = prog[prog[addend1pos]] + prog[prog[addend2pos]]

		case 2: // opcode `2` means multiplication.
			factor1pos := ip + 1
			factor2pos := ip + 2
			productPos := ip + 3
			prog[prog[productPos]] = prog[prog[factor1pos]] * prog[prog[factor2pos]]

		case 99: // opcode `99` means halt.
			break evaluation // End program evaluation.

		default: // Unknown opcode means we don't know what to do.
			log.Fatal("Unknown opcode:", prog[ip])
		}
	}

    // Return the program's final output/result from its address 0.
    return prog[0]
}

func main () {
    desiredResult := 19690720

    // Supply incrementing nouns and verbs until we get the result we
    // are seeking.
    for noun := 0; noun < 100; noun++ {
        for verb := 0; verb < 100; verb++ {
            result := run(noun, verb)
            if result == desiredResult {
                fmt.Println("Success found with noun", noun, "and verb", verb)
            }
        }
    }

}
