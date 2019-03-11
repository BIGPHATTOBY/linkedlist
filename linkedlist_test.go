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

func TestAppend(t *testing.T) {
	ll := NewLinkedList()
	ll.Append("foo")
	length, err := ll.Length()
	if err != nil {
		println(err)
	}
	if length != 1 {
		t.Error("Expected 1, got ", length)
	}
	ll.Append("bar")
	length, err = ll.Length()
	if err != nil {
		println(err)
	}
	if length != 2 {
		t.Error("Expected 2, got ", length)
	}
}

func TestInsert(t *testing.T) {
	ll := NewLinkedList()
	ll.Insert("foo", 0)
	length, err := ll.Length()
	if err != nil {
		t.Error(err)
	}
	if length != 1 {
		t.Error("Expected length to be 1, got ", length)
	}
	s, _ := ll.GetValueByIndex(1)
	if s != "foo" {
		t.Error("Expected \"foo\", got ", s)
	}
	ll.Insert("bar", 0)
	s, _ = ll.GetValueByIndex(1)
	if s != "bar" {
		t.Error("Expected \"bar\", got ", s)
	}
	s, _ = ll.GetValueByIndex(2)
	if s != "foo" {
		t.Error("Expected \"foo\", got ", s)
	}
}

func TestDelete(t *testing.T) {
	ll := NewLinkedList()
	ll.Append("foo")
	ll.Append("bar")
	length, err := ll.Length()
	if err != nil {
		t.Error(err)
	}
	if length != 2 {
		t.Error("Expected length to be 2, got ", length)
	}
	ll.Delete(1)
	length, err = ll.Length()
	if err != nil {
		t.Error(err)
	}
	if length != 1 {
		t.Error("Expected length to be 1, got ", length)
	}
}
