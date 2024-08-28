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
