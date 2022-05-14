package countPrice

import "testing"

func TestPrice(t *testing.T) {
	if value := price([]int{}); value != 0.0 {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{1}); value != 8.0 {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{2}); value != 8.0 {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{3}); value != 8.0 {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{4}); value != 8.0 {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{1, 1, 1}); value != 8.0*3 {
		t.Errorf("價格錯誤: %f", value)
	}
}

func TestPriceDiscounts(t *testing.T) {
	if value := price([]int{0, 1}); value != 8.0*2*0.95 {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{0, 2, 4}); value != 8.0*3*0.9 {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{0, 1, 2, 4}); value != 8.0*4*0.8 {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{0, 1, 2, 3, 4}); value != 8.0*5*0.75 {
		t.Errorf("價格錯誤: %f", value)
	}
}

func TestSeveralDiscounts(t *testing.T) {
	if value := price([]int{0, 0, 1}); value != 8+(8*2*0.95) {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{0, 0, 1, 1}); value != 2*(8*2*0.95) {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{0, 0, 1, 2, 2, 3}); value != (8*4*0.8)+(8*2*0.95) {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{0, 1, 1, 2, 3, 4}); value != 8+(8*5*0.75) {
		t.Errorf("價格錯誤: %f", value)
	}
}

func TestEdgeCases(t *testing.T) {
	if value := price([]int{0, 0, 1, 1, 2, 2, 3, 4}); value != 2*(8*4*0.8) {
		t.Errorf("價格錯誤: %f", value)
	}
	if value := price([]int{
		0, 0, 0, 0, 0,
		1, 1, 1, 1, 1,
		2, 2, 2, 2,
		3, 3, 3, 3, 3,
		4, 4, 4, 4,
	}); value != 3*(8*5*0.75)+2*(8*4*0.8) {
		t.Errorf("價格錯誤: %f", value)
	}
}
