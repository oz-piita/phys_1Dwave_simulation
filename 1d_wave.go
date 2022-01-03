package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	n       = 50  // length of system
	timemax = 100 // how many times repeat

	c  = 1.0  // velocity of wave
	dt = 0.12 // delta time
	dx = 0.1  // delta length of system
)

func main() {
	fy, ly, uy := createY(), createY(), createY()

	// to be stable, c should be c<=1.
	fmt.Println("c=", c*dt/dx)

	// calcurate with each scheme
	fy = ftcs(fy)
	ly = lax(ly)
	uy = upwind(uy)

	// save result in csv to plot with Python.
	outputCSV("./FTCSresult.csv", conv2d(fy))
	outputCSV("./LAXresult.csv", conv2d(ly))
	outputCSV("./UPWINDresult.csv", conv2d(uy))
}

func createY() [][]float64 {
	// Slice initialize with 0
	y := make([][]float64, timemax) // [time]*[x]
	for i := 0; i < timemax; i++ {
		k := make([]float64, n)
		y[i] = k
	}

	// initial condition
	for i := 0; i < n; i++ {
		a := float64(i) / float64(n)
		if i <= int(0.6*float64(n)) {
			y[0][i] = 1.25 * a
		} else {
			y[0][i] = 5.0 * (1.0 - a)
		}
	}
	return y
}

// FTCS(Forward in TIme and Central Differential in Space) scheme
func ftcs(y [][]float64) [][]float64 {
	for j := 0; j < timemax-1; j++ {
		for i := 1; i < n-1; i++ {
			y[j+1][i] = y[j][i] - (c*dt/2*dx)*(y[j][i+1]-y[j][i-1])
		}
	}
	return y
}

// Lax scheme
func lax(y [][]float64) [][]float64 {
	for j := 0; j < timemax-1; j++ {
		for i := 1; i < n-1; i++ {
			y[j+1][i] = (y[j][i+1]+y[j][i-1])/2 - (c*dt/2/dx)*(y[j][i+1]-y[j][i-1])
		}
	}
	return y
}

// 1st order UPWIND scheme
func upwind(y [][]float64) [][]float64 {
	for j := 0; j < timemax-1; j++ {
		for i := 1; i < n-1; i++ {
			y[j+1][i] = y[j][i] - (c*dt/dx)*(y[j][i]-y[j][i-1])
		}
	}
	return y
}

// 吐き出し時の浮動小数のフォーマットを固定してstringで返す
func conv2d(y [][]float64) [][]string {
	records := make([][]string, len(y))
	for i := range y {
		r := make([]string, len(y[0]))
		for j := range y[0] {
			r[j] = strconv.FormatFloat(y[i][j], 'f', -1, 64)
		}
		records[i] = r
	}
	return records
}

// create csvfile
func outputCSV(filepath string, records [][]string) {

	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	w.WriteAll(records)
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
