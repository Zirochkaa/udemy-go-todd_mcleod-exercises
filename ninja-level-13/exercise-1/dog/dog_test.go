package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	expected := 14
	got := Years(2)
	if got != expected {
		t.Error("Expected", expected, "Got", got)
	}
}

func TestYearsTwo(t *testing.T) {
	expected := 7
	got := YearsTwo(1)
	if got != expected {
		t.Error("Expected", expected, "Got", got)
	}
}

func ExampleYears() {
	fmt.Println(Years(2))
	// Output:
	// 14
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(1))
	// Output:
	// 7
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(6)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
