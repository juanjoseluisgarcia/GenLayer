package main

import (
	"githuhub.com/juanjoseluisgarcia/GenLayer/cmd/01-data-fragmentation/model"
	"strconv"
)

type HashDataHandler struct{}

func (hdh *HashDataHandler) ReconstructData(fragments []model.Fragment) string {
	var reconstructedData string
	for _, fragment := range fragments {
		if fragment.Hash != SimpleHash(fragment.Data) {
			return ""
		}
		reconstructedData += fragment.Data
	}

	return reconstructedData
}

func SimpleHash(data string) string {
	var hashValue int64 = 0
	mod := int64(1e9 + 7)

	for i, char := range data {
		hashValue += int64(char) * int64(i+1)
		hashValue %= mod
	}

	hashStr := strconv.FormatInt(hashValue, 36)
	for len(hashStr) < 30 {
		hashStr += "="
	}

	return hashStr[:30]
}
