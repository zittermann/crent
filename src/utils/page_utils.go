package utils

import "math"

func GetTotalPages(totalElements int64, limit int) int {
	/* Calculamos la cantidad de paginas */
	return int(math.Ceil(
		float64(totalElements) / float64(limit)))
}