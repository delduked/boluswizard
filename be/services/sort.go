package services

import (
	"api/types"
	"sort"
)

// Sort list of corrections from lowest to highest
func SortCorrection(value []types.Correction) {
	sort.SliceStable(value, func(i, j int) bool {
		return value[i].TimeStamp < value[j].TimeStamp
	})
}
