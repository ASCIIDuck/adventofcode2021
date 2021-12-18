package main

import (
	"testing"
)

func TestCalculateNestedAdds(t *testing.T) {
	test_data := Packet{packet_type: 0, sub_packets: []Packet{
		{packet_type: 0, sub_packets: []Packet{
			{packet_type: 0, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 0, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		}},
		{packet_type: 0, sub_packets: []Packet{
			{packet_type: 0, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 0, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		},
		}}}
	expected_sum := 20
	if sum := calculate(test_data); sum != expected_sum {
		t.Errorf("calculate() = %d, want %d", sum, expected_sum)
	}
}

func TestCalculateNestedMultiply(t *testing.T) {
	test_data := Packet{packet_type: 1, sub_packets: []Packet{
		{packet_type: 1, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		}},
		{packet_type: 1, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		},
		}}}
	expected_result := 576
	if result := calculate(test_data); result != expected_result {
		t.Errorf("calculate() = %d, want %d", result, expected_result)
	}
}

func TestCalculateNestedMin(t *testing.T) {
	test_data := Packet{packet_type: 2, sub_packets: []Packet{
		{packet_type: 2, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 0, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		}},
		{packet_type: 2, sub_packets: []Packet{
			{packet_type: 0, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		},
		}}}
	expected_result := 2
	if result := calculate(test_data); result != expected_result {
		t.Errorf("calculate() = %d, want %d", result, expected_result)
	}
}

func TestCalculateNestedMax(t *testing.T) {
	test_data := Packet{packet_type: 3, sub_packets: []Packet{
		{packet_type: 3, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 0, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		}},
		{packet_type: 3, sub_packets: []Packet{
			{packet_type: 0, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		},
		}}}
	expected_product := 12
	if product := calculate(test_data); product != expected_product {
		t.Errorf("calculate() = %d, want %d", product, expected_product)
	}
}

func TestCalculateNestedWithGreaterThanExpectFalse(t *testing.T) {
	test_data := Packet{packet_type: 5, sub_packets: []Packet{
		{packet_type: 0, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		}},
		{packet_type: 1, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		},
		}}}
	expected_result := 0
	if result := calculate(test_data); result != expected_result {
		t.Errorf("calculate() = %d, want %d", result, expected_result)
	}
}

func TestCalculateNestedWithGreaterThanExceptTrue(t *testing.T) {
	test_data := Packet{packet_type: 5, sub_packets: []Packet{
		{packet_type: 1, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		}},
		{packet_type: 0, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		},
		}}}
	expected_result := 1
	if result := calculate(test_data); result != expected_result {
		t.Errorf("calculate() = %d, want %d", result, expected_result)
	}
}

func TestCalculateNestedWithLessThanExpectTrue(t *testing.T) {
	test_data := Packet{packet_type: 6, sub_packets: []Packet{
		{packet_type: 0, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		}},
		{packet_type: 1, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		},
		}}}
	expected_result := 1
	if result := calculate(test_data); result != expected_result {
		t.Errorf("calculate() = %d, want %d", result, expected_result)
	}
}

func TestCalculateNestedWithLessThanExpectFalse(t *testing.T) {
	test_data := Packet{packet_type: 6, sub_packets: []Packet{
		{packet_type: 1, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		}},
		{packet_type: 0, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		},
		}}}
	expected_result := 0
	if result := calculate(test_data); result != expected_result {
		t.Errorf("calculate() = %d, want %d", result, expected_result)
	}
}

func TestCalculateEuqalExpectFalse(t *testing.T) {
	test_data := Packet{packet_type: 7, sub_packets: []Packet{
		{packet_type: 4, literal: 4},
		{packet_type: 4, literal: 3},
	}}

	expected_result := 0
	if result := calculate(test_data); result != expected_result {
		t.Errorf("calculate() = %d, want %d", result, expected_result)
	}
}

func TestCalculateEuqalExpectTrue(t *testing.T) {
	test_data := Packet{packet_type: 7, sub_packets: []Packet{
		{packet_type: 4, literal: 3},
		{packet_type: 4, literal: 3},
	}}

	expected_result := 1
	if result := calculate(test_data); result != expected_result {
		t.Errorf("calculate() = %d, want %d", result, expected_result)
	}
}
func TestCalculateNestedWithEqualExpectFalse(t *testing.T) {
	test_data := Packet{packet_type: 7, sub_packets: []Packet{
		{packet_type: 1, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		}},
		{packet_type: 0, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		},
		}}}
	expected_result := 0
	if result := calculate(test_data); result != expected_result {
		t.Errorf("calculate() = %d, want %d", result, expected_result)
	}
}

func TestCalculateNestedWithEqualExpectTrue(t *testing.T) {
	test_data := Packet{packet_type: 7, sub_packets: []Packet{
		{packet_type: 0, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		}},
		{packet_type: 0, sub_packets: []Packet{
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 1}, {packet_type: 4, literal: 2}}},
			{packet_type: 1, sub_packets: []Packet{{packet_type: 4, literal: 3}, {packet_type: 4, literal: 4}}},
		},
		}}}
	expected_result := 1
	if result := calculate(test_data); result != expected_result {
		t.Errorf("calculate() = %d, want %d", result, expected_result)
	}
}

func TestReadLiteral2021(t *testing.T) {
	input := "101111111000101"
	expected_result := 2021
	expected_remaining := ""
	if result, remaining := read_literal(input); result != expected_result || remaining != expected_remaining {
		t.Errorf("read_literal() = (%d, %s), want (%d,%s)", result, remaining, expected_result, expected_remaining)
	}
}
func TestReadLiteral2021WithExtra(t *testing.T) {
	input := "101111111000101000"
	expected_result := 2021
	expected_remaining := "000"
	if result, remaining := read_literal(input); result != expected_result || remaining != expected_remaining {
		t.Errorf("read_literal() = (%d, %s), want (%d,%s)", result, remaining, expected_result, expected_remaining)
	}
}

func TestReadLiteral5(t *testing.T) {
	input := "001010000"
	expected_result := 5
	expected_remaining := "0000"
	if result, remaining := read_literal(input); result != expected_result || remaining != expected_remaining {
		t.Errorf("read_literal() = (%d, %s), want (%d,%s)", result, remaining, expected_result, expected_remaining)
	}
}
