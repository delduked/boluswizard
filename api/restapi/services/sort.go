package services

import (
	"boluswizard/models"
	"sort"
)

// Sort list of corrections from lowest to highest
func SortCorrection(value []*models.Correction) {
	sort.SliceStable(value, func(i, j int) bool {
		return value[i].TimeStamp < value[j].TimeStamp
	})
}
