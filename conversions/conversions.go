package conversions

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for stringIndex, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {

			return nil, errors.New("failed to convert")
		}

		floats[stringIndex] = floatVal
	}
	return floats, nil
}
