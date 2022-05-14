package countPrice

func price(cart []int) float64 {
	if len(cart) == 0 {
		return 0.0
	}
	if len(cart) == 1 {
		return 8.0
	}
	cS := [5]int{0, 0, 0, 0, 0}
	for i := range cart {
		cS[cart[i]] += 1
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
	return P
}
