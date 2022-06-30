package statistic

import "github.com/montanaflynn/stats"

func Mean_Function(Data []float64) float64 {
	var Mean_answer float64

	Mean_answer, _ = stats.Mean(Data)
	return Mean_answer
}
func Median_function(Data []float64) float64 {
	var Median_answer float64
	Median_answer, _ = stats.Median(Data)
	return Median_answer
}
func Mode_function(Data []float64) float64 {
	var Mode_answer float64
	Mode_answer, _ = stats.Median(Data)
	return Mode_answer
}
