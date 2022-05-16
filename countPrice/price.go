package countPrice

import (
	"strconv"
)

var store = make(map[string]float64)

func Price(cart []int) float64 {

	if len(cart) == 0 {
		return 0.0
	}
	if len(cart) == 1 {
		return 8.0
	}
	cS := []int{0, 0, 0, 0, 0}
	for _, i := range cart {
		cS[i] += 1
	}
	return priceConvert(cS)
}

func priceConvert(cS []int) float64 {
	val, exist := store[strconv.Itoa(cS[0])+strconv.Itoa(cS[1])+strconv.Itoa(cS[2])+strconv.Itoa(cS[3])+strconv.Itoa(cS[4])]
	if exist {
		return val
	}
	f := false
	for _, j := range cS {
		if j > 1 {
			f = true
			break
		}
	}
	if f {
		var storeList []float64
		l32 := [][]int{
			{0, 0, 0, 0, 1},
			{1, 0, 0, 0, 0},
			{1, 0, 0, 0, 1},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 1, 1},
			{1, 0, 0, 1, 0},
			{1, 0, 0, 1, 1},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 1},
			{1, 0, 1, 0, 0},
			{1, 0, 1, 0, 1},
			{0, 0, 1, 1, 0},
			{0, 0, 1, 1, 1},
			{1, 0, 1, 1, 0},
			{1, 0, 1, 1, 1},
			{0, 1, 0, 0, 0},
			{0, 1, 0, 0, 1},
			{1, 1, 0, 0, 0},
			{1, 1, 0, 0, 1},
			{0, 1, 0, 1, 0},
			{0, 1, 0, 1, 1},
			{1, 1, 0, 1, 0},
			{1, 1, 0, 1, 1},
			{0, 1, 1, 0, 0},
			{0, 1, 1, 0, 1},
			{1, 1, 1, 0, 0},
			{1, 1, 1, 0, 1},
			{0, 1, 1, 1, 0},
			{0, 1, 1, 1, 1},
			{1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1},
		}
		for _, v := range l32 {
			if cS[0] >= v[0] && cS[1] >= v[1] && cS[2] >= v[2] && cS[3] >= v[3] && cS[4] >= v[4] {
				vl := []int{cS[0] - v[0], cS[1] - v[1], cS[2] - v[2], cS[3] - v[3], cS[4] - v[4]}
				t := priceConvert(vl)
				store[l2s(vl)] = t
				storeList = append(storeList, t+priceConvert(v))
			}
		}
		min := 9999.0
		for _, i := range storeList {
			if i < min {
				min = i
			}
		}
		return min
	} else {
		return priceSimple(cS)
	}

}
func l2s(cS []int) string {
	return strconv.Itoa(cS[0]) + strconv.Itoa(cS[1]) + strconv.Itoa(cS[2]) + strconv.Itoa(cS[3]) + strconv.Itoa(cS[4])
}
func priceSimple(cS []int) float64 {
	val, exist := store[l2s(cS)]
	if exist {
		return val
	}
	for _, i := range cS {
		if i < 0 {
			return 999.0
		}
	}
	P := 0.0
	for {
		if cS[0] > 0 && cS[1] > 0 && cS[2] > 0 && cS[3] > 0 && cS[4] > 0 {
			for i := 0; i < 5; i++ {
				cS[i] -= 1
			}
			P += 40 * 0.75
		} else {
			break
		}
	}
	for {

		if cS[0] > 0 && cS[1] > 0 && cS[2] > 0 && cS[3] > 0 {
			t := []int{0, 1, 2, 3}
			for i := range t {
				cS[t[i]] -= 1
			}
			P += 32 * 0.8
		} else if cS[0] > 0 && cS[1] > 0 && cS[2] > 0 && cS[4] > 0 {
			t := []int{0, 1, 2, 4}
			for i := range t {
				cS[t[i]] -= 1
			}
			P += 32 * 0.8
		} else if cS[0] > 0 && cS[1] > 0 && cS[4] > 0 && cS[3] > 0 {
			t := []int{0, 1, 4, 3}
			for i := range t {
				cS[t[i]] -= 1
			}
			P += 32 * 0.8
		} else if cS[0] > 0 && cS[4] > 0 && cS[2] > 0 && cS[3] > 0 {
			t := []int{0, 4, 2, 3}
			for i := range t {
				cS[t[i]] -= 1
			}
			P += 32 * 0.8
		} else if cS[4] > 0 && cS[1] > 0 && cS[2] > 0 && cS[3] > 0 {
			t := []int{0, 4, 2, 3}
			for i := range t {
				cS[t[i]] -= 1
			}
			P += 32 * 0.8
		} else {
			break
		}
	}
	c := [][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 1, 4},
		{0, 2, 3},
		{0, 2, 4},
		{0, 3, 4},
		{1, 2, 3},
		{1, 2, 4},
		{1, 3, 4},
		{2, 3, 4},
		{0, 1, 2},
	}
	for {
		f := false
	r:
		for i := range c {
			if cS[c[i][0]] > 0 && cS[c[i][1]] > 0 && cS[c[i][2]] > 0 {
				for j := range c[i] {
					cS[c[i][j]] -= 1
					f = true
				}
				P += 24 * 0.9
			}
		}
		if f {
			f = false
			goto r
		} else {
			break
		}
	}
	c = [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 2},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	for {
		f := false
	s:
		for i := range c {
			if cS[c[i][0]] > 0 && cS[c[i][1]] > 0 {
				for j := range c[i] {
					cS[c[i][j]] -= 1
					f = true
				}
				P += 16 * 0.95
			}
		}
		if f {
			f = false
			goto s
		} else {
			break
		}
	}
	for _, i := range cS {
		P += float64(8 * i)
	}
	store[strconv.Itoa(cS[0])+strconv.Itoa(cS[1])+strconv.Itoa(cS[2])+strconv.Itoa(cS[3])+strconv.Itoa(cS[4])] = P
	return P
}
