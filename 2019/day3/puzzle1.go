package main

// --- Day 3: Crossed Wires ---
//
// The gravity assist was successful, and you're well on your way to the Venus refuelling station. During the rush back on Earth, the fuel management system wasn't completely installed, so that's next on the priority list.
//
// Opening the front panel reveals a jumble of wires. Specifically, two wires are connected to a central port and extend outward on a grid. You trace the path each wire takes as it leaves the central port, one wire per line of text (your puzzle input).
//
// The wires twist and turn, but the two wires occasionally cross paths. To fix the circuit, you need to find the intersection point closest to the central port. Because the wires are on a grid, use the Manhattan distance for this measurement. While the wires do technically cross right at the central port where they both start, this point does not count, nor does a wire count as crossing with itself.
//
// For example, if the first wire's path is R8,U5,L5,D3, then starting from the central port (o), it goes right 8, up 5, left 5, and finally down 3:
//
// ...........
// ...........
// ...........
// ....+----+.
// ....|....|.
// ....|....|.
// ....|....|.
// .........|.
// .o-------+.
// ...........
//
// Then, if the second wire's path is U7,R6,D4,L4, it goes up 7, right 6, down 4, and left 4:
//
// ...........
// .+-----+...
// .|.....|...
// .|..+--X-+.
// .|..|..|.|.
// .|.-X--+.|.
// .|..|....|.
// .|.......|.
// .o-------+.
// ...........
//
// These wires cross at two locations (marked X), but the lower-left one is closer to the central port: its distance is 3 + 3 = 6.
//
// Here are a few more examples:
//
// - R75,D30,R83,U83,L12,D49,R71,U7,L72
// U62,R66,U55,R34,D71,R55,D58,R83 = distance 159
// - R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
// U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135
//
// What is the Manhattan distance from the central port to the closest intersection?

import (
	"encoding/csv"
	"fmt"
	intersect "github.com/juliangruber/go-intersect"
	"log"
	"math"
	"os"
	"strconv"
)

// traceWirePositions returns slice of two-element integer arrays that
// correspond to the X and Y coordinates of the path that a given wire
// takes across a plane.
func traceWirePositions(wire []string) [][2]int {
	var x int
	var y int
	var positions [][2]int

	for _, v := range wire {
		distance, _ := strconv.Atoi(v[1:])
		for i := 0; i < distance; i++ {
			direction := v[0]
			switch direction {
			case 'R':
				x += 1
			case 'L':
				x -= 1
			case 'D':
				y += 1
			case 'U':
				y -= 1
			}

			positions = append(positions, [2]int{x, y})
		}
	}

	return positions
}

// manhattanDistance implements the simple taxicab geometry arithmetic
// and returns the total distance of both the X and Y axes travelled.
func manhattanDistance(coords [2]int) float64 {
	return math.Abs(float64(coords[0])) + math.Abs(float64(coords[1]))
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	inputs, _ := r.ReadAll()

	wire1 := inputs[0]
	wire2 := inputs[1]

	// Read in the directional instructions and create a trace that
	// records every coordinate that would be touched by that wire.
	trace1 := traceWirePositions(wire1)
	trace2 := traceWirePositions(wire2)

	// Find the locations where both traces overlap.
	x := intersect.Hash(trace1, trace2)

	// Get the Manhattan distance values for each instance of overlap.
	var distances []float64
	for _, v := range x {
		distances = append(distances, manhattanDistance(v.([2]int)))
	}

	// Find the minimum of these distance values.
	z := math.Inf(0)
	for _, v := range distances {
		z = (math.Min(z, v))
	}
	fmt.Println("Minimum distance is:", z)
}
