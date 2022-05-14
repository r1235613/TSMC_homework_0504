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
