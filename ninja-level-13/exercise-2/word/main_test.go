package word

import (
	"fmt"
	"ninjalevel12/exercise2/quote"
	"testing"
)

func TestCount(t *testing.T) {
	expected := 3
	got := Count("some string example")
	if got != expected {
		t.Error("Expected", expected, "Got", got)
	}
}

func TestUseCount(t *testing.T) {
	got := UseCount("some string example string example")
	for k, v := range got {
		switch k {
		case "some":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		case "string":
			if v != 2 {
				t.Error("Expected", 2, "Got", v)
			}
		case "example":
			if v != 2 {
				t.Error("Expected", 2, "Got", v)
			}
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("some string example"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
