package xlib

import "sort"

// MapStringFloat for sort map[string]int
type MapStringFloat map[string]float64

type pairStringFloat struct {
	Key   string
	Value float64
}

// PairStringFloatSlice is []struct{string,float64}
type PairStringFloatSlice []pairStringFloat

func (p PairStringFloatSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairStringFloatSlice) Len() int           { return len(p) }
func (p PairStringFloatSlice) Less(i, j int) bool { return p[i].Value < p[j].Value }

// Sort method of MapStringFloat
func (m *MapStringFloat) Sort(ascend bool) PairStringFloatSlice {
	p := make(PairStringFloatSlice, len(*m))
	i := 0
	for k, v := range *m {
		p[i] = pairStringFloat{k, v}
		i++
	}
	if ascend {
		sort.Sort(p)
	} else {
		sort.Sort(sort.Reverse(p))
	}
	return p
}

// MapStringInt for sort map[string]int
type MapStringInt map[string]int

// TODO:

type pairStringInt struct {
	Key   string
	Value int
}

// PairStringIntSlice is []struct{string, int}
type PairStringIntSlice []pairStringInt

func (p PairStringIntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairStringIntSlice) Len() int           { return len(p) }
func (p PairStringIntSlice) Less(i, j int) bool { return p[i].Value < p[j].Value }

// Sort method of MapStringInt
func (m MapStringInt) Sort(ascend bool) PairStringIntSlice {
	p := make(PairStringIntSlice, len(m))
	i := 0
	for k, v := range m {
		p[i] = pairStringInt{k, v}
		i++
	}
	if ascend {
		sort.Sort(p)
	} else {
		sort.Sort(sort.Reverse(p))
	}
	return p
}

func getKeysInt(m map[int]bool) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func getKeysString(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
