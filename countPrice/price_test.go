package countPrice

import "testing"

func TestPrice(t *testing.T) {
	if value := Price([]int{}); value != 0.0 {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8.0)
	}
	if value := Price([]int{1}); value != 8.0 {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8.0)
	}
	if value := Price([]int{2}); value != 8.0 {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8.0)
	}
	if value := Price([]int{3}); value != 8.0 {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8.0)
	}
	if value := Price([]int{4}); value != 8.0 {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8.0)
	}
	if value := Price([]int{1, 1, 1}); value != 8.0*3 {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8.0*3)
	}
}

func TestPriceDiscounts(t *testing.T) {
	if value := Price([]int{0, 1}); value != 8.0*2*0.95 {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8.0*2*0.95)
	}
	if value := Price([]int{0, 2, 4}); value != 8.0*3*0.9 {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8.0*3*0.9)
	}
	if value := Price([]int{0, 1, 2, 4}); value != 8.0*4*0.8 {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8.0*4*0.8)
	}
	if value := Price([]int{0, 1, 2, 3, 4}); value != 8.0*5*0.75 {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8.0*5*0.75)
	}
}

func TestSeveralDiscounts1(t *testing.T) {
	if value := Price([]int{0, 0, 1}); value != 8+(8*2*0.95) {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8+(8*2*0.95))
	}
}
func TestSeveralDiscounts2(t *testing.T) {
	if value := Price([]int{0, 0, 1, 1}); value != 2*(8*2*0.95) {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 2*(8*2*0.95))
	}
}
func TestSeveralDiscounts3(t *testing.T) {
	if value := Price([]int{0, 0, 1, 2, 2, 3}); value != (8*4*0.8)+(8*2*0.95) {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, (8*4*0.8)+(8*2*0.95))
	}
}
func TestSeveralDiscounts4(t *testing.T) {
	if value := Price([]int{0, 1, 1, 2, 3, 4}); value != 8+(8*5*0.75) {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 8+(8*5*0.75))
	}
}

func TestEdgeCases1(t *testing.T) {
	if value := Price([]int{0, 0, 1, 1, 2, 2, 3, 4}); value != 2*(8*4*0.8) {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 2*(8*4*0.8))
	}
}

func TestEdgeCases2(t *testing.T) {
	if value := Price([]int{
		0, 0, 0, 0, 0,
		1, 1, 1, 1, 1,
		2, 2, 2, 2,
		3, 3, 3, 3, 3,
		4, 4, 4, 4,
	}); value != 3*(8*5*0.75)+2*(8*4*0.8) {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 3*(8*5*0.75)+2*(8*4*0.8))
	}

}

func TestCasesFromWeb1(t *testing.T) {
	if value := Price([]int{0, 1, 1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4}); value != 100 {
		t.Errorf("價格錯誤: %f，正確應為%f元", value, 100.0)
	}

}
