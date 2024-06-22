package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64

	for _, strValue := range strings {
		floatValue, err := strconv.ParseFloat(strValue, 64)

		if err != nil {
			return nil, errors.New("Failed converting strings to floats")
		}

		floats = append(floats, floatValue)
	}

	return floats, nil
}
