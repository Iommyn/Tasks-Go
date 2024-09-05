package Task

import (
	"reflect"
	"testing"
)

func TestTaskOne(t *testing.T) {
	input := "hello world is a epta"
	want := map[string]int{"hello": 1, "world": 1, "epta": 1}
	got := TaskOne(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskOne(%q) = %v; want %v", input, got, want)
	}
}

func TestTaskTwo(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	want := 55
	got := TaskTwo(input)
	if got != want {
		t.Errorf("TaskTwo(%v) = %d; want %d", input, got, want)
	}
}

func TestTaskThree(t *testing.T) {
	map1 := map[string]int{"one": 5, "two": 15}
	map2 := map[string]int{"two": 6, "three": 12}
	want := map[string]int{"one": 5, "two": 21, "three": 12}
	got := TaskThree(map1, map2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskThree(%v, %v) = %v; want %v", map1, map2, got, want)
	}
}

func TestTaskFour(t *testing.T) {
	tests := []struct {
		input string
		want  map[string]int
	}{
		{"Hello, world! Hello, Go!", map[string]int{"hello": 2, "world": 1, "go": 1}},
		{"This is a test. This test is only a test.", map[string]int{"this": 2, "is": 2, "a": 2, "test": 3, "only": 1}},
	}

	for _, test := range tests {
		got := TaskFour(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("TaskFour(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestTaskFive(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{15}},
		{[]int{1, 1, 1, 1, 1}, []int{5}},
		{[]int{5, 4, 3, 2, 1}, []int{15}},
	}

	for _, test := range tests {
		got := TaskFive(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("TaskFive(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}
func TestTaskSix(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
	}

	for _, test := range tests {
		if got := TaskSix(test.input); got != test.want {
			t.Errorf("TaskSix(%q) = %q; want %q", test.input, got, test.want)
		}
	}
}

func TestTaskSeven(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"apple banana cherry date", []string{"apple", "banana", "cherry", "date"}},
	}

	for _, test := range tests {
		if got := TaskSeven(test.input); !equal(got, test.want) {
			t.Errorf("TaskSeven(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
