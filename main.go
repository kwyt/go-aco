package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"strings"
)

func main() {
	an := flag.Int("a", 500, "The number of agent")
	tc := flag.Int("t", 300, "Number of trials")
	ex := flag.Bool("export", false, "Result by plotting the output")
	flag.Parse()

	graph := NewGraph()
	graph.initialize(load())
	aco := NewAco()
	aco.run(graph, *an, *tc, *ex)
}

func load() [][]float64 {

	file := OpenFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var coords = [][]float64{}

	for scanner.Scan() {
		line := scanner.Text()
		coords = append(coords, stringToFloat(strings.Split(line, " ")))
	}
	return coords
}

func stringToFloat(s []string) []float64 {
	floatArr := make([]float64, len(s))

	for i := 0; i < len(s); i++ {
		floatArr[i], _ = strconv.ParseFloat(s[i], 64)
	}

	return floatArr
}

func OpenFile() *os.File {
	fp, err := os.Open("./sample.dat")

	if err != nil {
		panic(err)
	}

	return fp
}
