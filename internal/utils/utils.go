package utils

import "errors"

func Split(slice []int, size int) ([][]int, error) {
	if slice == nil || size <= 0 {
		return nil, errors.New("Illegal argument error")
	}

	var resultSize = (len(slice) + size - 1) / size
	var result = make([][]int, resultSize)
	for i := 0; i < resultSize; i++ {
		var right = size * i
		var left = size * (i + 1)
		if left > len(slice) {
			left = len(slice)
		}
		result[i] = slice[right:left]
	}
	return result, nil
}

func RevertMap(sourceMap map[string]string) (map[string]string, error) {
	if sourceMap == nil {
		return nil, errors.New("Illegal argument error")
	}

	var result map[string]string = make(map[string]string, len(sourceMap))
	for key, value := range sourceMap {
		result[value] = key
	}
	return result, nil
}

var hardcodedSet = map[string]bool{"value1": true, "value2": true}

func Filter(slice []string) ([]string, error) {
	if slice == nil {
		return nil, errors.New("Illegal argument error")
	}
	var result = make([]string, 0)
	for i := 0; i < len(slice); i++ {
		if hardcodedSet[slice[i]] == true {
			result = append(result, slice[i])
		}
	}
	return result, nil
}
