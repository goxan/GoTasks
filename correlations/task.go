ackage main

import (
	"encoding/csv"
	"log"
	"os"
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

func pearsonCorrelation(a, b []float32) float32 {
	//TODO replace
	return 0
}

func spearmanCorrelation(a, b []float32) float32 {
	//TODO replace
	return 0
}
func convertToNumbers(rows [][]string) ([][]float32, error) {
	//TODO replace
	return make([][]float32, 0), nil
}

func handleError(e error, msg string) {
	if e != nil {
		log.Fatal(msg)
	}
}

func rotatedDataset(x [][]float32) [][]float32 {
	//TODO replace
	return x
}

func main() {
	rows := readCsv("housing.csv")
	rowWiseDatset, error := convertToNumbers(rows)
	columnWiseDataset := rotatedDataset(rowWiseDatset)
	if error != nil {
		// handle
	}
	// Do your analysis here, identify what affects the house price
	pearsonCorrelation(columnWiseDataset[0], columnWiseDataset[1]) //

}
