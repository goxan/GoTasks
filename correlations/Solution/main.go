package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func readCsv(path string) [][]string {
	file, err := os.Open(path)
	errorReadingProblem := "The file does not exists. Please check the file path"
	handleError(err, errorReadingProblem)
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	handleError(err, errorReadingProblem)
	return records
}

func pearsonCorrelation(a, b []float64) float64 {

	if len(a) != len(b) || len(a) == 0 {
		return 0
	}
	var A, B float64
	for i := 0; i < len(a); i++ {
		A += a[i]
		B += b[i]
	}
	var resA = A / float64(len(a)) //средние а и б
	var resB = B / float64(len(b))

	var covariance float64
	for i := 0; i < len(a); i++ {
		covariance += float64((a[i] - resA) * (b[i] - resB))
	}
	covariance /= float64(len(a))

	var deviationA float64
	var deviationB float64

	for i := 0; i < len(a); i++ {
		deviationA += math.Pow(float64(a[i]-resA), 2)
		deviationB += math.Pow(float64(b[i]-resB), 2)
	}

	var newResA = math.Sqrt(deviationA) / float64(len(a))
	var newResB = math.Sqrt(deviationB) / float64(len(b))

	deviationA = newResA
	deviationB = newResB

	var PearsonCorrelationCoefficient float64

	PearsonCorrelationCoefficient = covariance / (deviationA * deviationB)

	return PearsonCorrelationCoefficient
}

func rank(value []float64) []float64 {
	type Pair struct {
		Value float64
		Index int
	}
	pairs := make([]Pair, len(value))
	for i, v := range value {
		pairs[i] = Pair{Value: v, Index: i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	ranks := make([]float64, len(value))
	for i, pair := range pairs {
		ranks[pair.Index] = float64(i + 1)
	}
	return ranks
}

func spearmanCorrelation(a, b []float64) float64 {

	rankA := rank(a)
	rankB := rank(b)

	var squareDifferenceSum float64
	for i := 0; i < len(a); i++ {
		diff := rankA[i] - rankB[i]
		squareDiff := diff * diff
		squareDifferenceSum += squareDiff
	}

	n := float64(len(a))
	spearman := 1 - (6*squareDifferenceSum)/(n*(n*n-1))

	return spearman
}
func convertToNumbers(rows [][]string) ([][]float64, error) {
	var result [][]float64
	for _, row := range rows {
		var convertedRow []float64
		for _, cell := range row {
			value, err := strconv.ParseFloat(cell, 64)
			if err != nil {
				convertedRow = append(convertedRow, 0.0)
			}
			convertedRow = append(convertedRow, float64(value))
		}
		result = append(result, convertedRow)
	}
	/*	fmt.Println(result)*/
	return result, nil
}

func handleError(e error, msg string) {
	if e != nil {
		log.Fatal(msg)
	}
}

func rotatedDataset(x [][]float64) [][]float64 {
	if len(x) == 0 || len(x[0]) == 0 {
		// Обработка случая, когда x пустой или вложенные срезы пусты
		return nil
	}
	firstRowLen := len(x[0])

	newArr := make([][]float64, firstRowLen)
	for i := 0; i < firstRowLen; i++ {
		for j := 0; j < len(x[0]); j++ {
			newArr[j] = append(newArr[j], x[i][j])
		}
	}

	return newArr
}

func main() {
	rows := readCsv("housing.csv")

	floatRows, _ := convertToNumbers(rows)
	floatRows = floatRows[1:]

	res := rotatedDataset(floatRows)

	var PearsonCorrelationCoefficient = pearsonCorrelation(res[2], res[8])

	var SpearmanCorrelationCoefficient = spearmanCorrelation(res[2], res[8])

	fmt.Println(SpearmanCorrelationCoefficient)
	fmt.Println(PearsonCorrelationCoefficient)

}
