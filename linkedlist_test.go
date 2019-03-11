package linkedlist

import (
	"testing"
)

func TestLength(t *testing.T) {
	ll := NewLinkedList()
	ll.Append("foo")
	ll.Append("bar")
	ll.Append("baz")
	length, err := ll.Length()
	if err != nil {
		println(err)
	}
	if length != 3 {
		t.Error("Expected 3, got ", length)
	}
}

func TestPrintAll(t *testing.T) {
	ll := NewLinkedList()
	ll.Append("foo")
	ll.Append("bar")
	ll.Append("baz")
	s, _ := ll.PrintAll()
	for i, v := range s {
		if i == 0 {
			if v != "foo" {
				t.Error("Expected foo, at index", i)
			}
		}
		if i == 1 {
			if v != "bar" {
				t.Error("Expected bar, at index", i)
			}
		}
		if i == 2 {
			if v != "baz" {
				t.Error("Expected baz, at index", i)
			}
		}
	}
}

func TestGetValueByIndex(t *testing.T) {
	ll := NewLinkedList()
	ll.Append("foo")
	ll.Append("bar")
	ll.Append("baz")
	s, _ := ll.GetValueByIndex(2)
	if s != "bar" {
		t.Error("Expected bar, got ", s)
	}
	s, _ = ll.GetValueByIndex(3)
	if s != "baz" {
		t.Error("Expected baz, got ", s)
	}
}
