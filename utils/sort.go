package utils

import (
	"sort"
)

func revSortByValue(mp map[string]float64) PairList {
	pl := make(PairList, len(mp))
	i := 0
	for k, v := range mp {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

// Pair struct
type Pair struct {
	Key   string
	Value float64
}

// PairList type
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
