package strings

import (
	"testing"
)

func TestLower(t *testing.T) {
	result := Lower("ThIS is MY NaMe")
	expected := "this is my name"

	if result != expected {
		t.Errorf("Expected \"%s\" got \"%s\"", expected, result)
	}
}

func TestUpper(t *testing.T) {
	result := Upper("ThIS is MY NaMe")
	expected := "THIS IS MY NAME"

	if result != expected {
		t.Errorf("Expected \"%s\" got \"%s\"", expected, result)
	}
}

func TestAfter(t *testing.T) {
	result := After("this is my name", "this ")
	expected := "is my name"

	if result != expected {
		t.Errorf("Expected \"%s\" got \"%s\"", expected, result)
	}

	result = After("this is my name", "nothing to match")
	if result != "" {
		t.Error("Calling After with nothing that matches should return empty string")
	}
}

func TestBefore(t *testing.T) {
	result := Before("this is my name", " is my name")
	expected := "this"

	if result != expected {
		t.Errorf("Expected \"%s\" got \"%s\"", expected, result)
	}

	result = Before("this is my name", "nothing to match")
	if result != "this is my name" {
		t.Error("Calling Before with nothing that matches should return original string")
	}
}

func TestContains(t *testing.T) {
	result := Contains("this is my name", "name")
	if !result {
		t.Errorf("Expected contains to find name in original string")
	}

	result = Contains("this is my name", "hello")
	if result {
		t.Errorf("Expected contains to not find hello in original string")
	}
}

func TestContainsSome(t *testing.T) {
	result := ContainsSome("this is my name", []string{"name", "hello"})
	if !result {
		t.Errorf("Expected contains to find name in original string")
	}

	result = ContainsSome("this is my name", []string{"hello", "world"})
	if result {
		t.Errorf("Expected contains to not find hello or world in original string")
	}
}

func TestContainsAll(t *testing.T) {
	result := ContainsAll("this is my name", []string{"name", "this"})
	if !result {
		t.Errorf("Expected contains to find name and this in original string")
	}

	result = ContainsAll("this is my name", []string{"name", "hello"})
	if result {
		t.Errorf("Expected contains to not find hello in original string")
	}
}

func TestLength(t *testing.T) {
	result := Length("This is my name")
	if result != 15 {
		t.Errorf("\"This is my name\" should have a length of 15 but got %d", result)
	}
}

func TestReverse(t *testing.T) {
	result := Reverse("This is my name")
	expected := "eman ym si sihT"

	if result != expected {
		t.Errorf("Reverse of This is my name should be eman ym si sihT, got %s", result)
	}
}
