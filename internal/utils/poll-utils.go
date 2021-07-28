package utils

import "errors"
import "github.com/ozoncp/ocp-poll-api/internal/models"


func SplitPoll(slice []models.Poll, size int) ([][]models.Poll, error) {
	if slice == nil || size <= 0 {
		return nil, errors.New("Illegal argument error")
	}

	var resultSize = (len(slice) + size - 1) / size
	var result = make([][]models.Poll, resultSize)
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

func PollsToMap(polls []models.Poll) (map[uint64]models.Poll, error) {
	if polls == nil {
		return nil, errors.New("Illegal argument error")
	}

	var result = make(map[uint64]models.Poll, len(polls))
	for _, poll := range polls {
		result[poll.Id] = poll
	}
	return result, nil
}

