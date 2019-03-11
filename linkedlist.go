package linkedlist

import (
	"errors"
)

type LinkedList struct {
	head *linkNode
}

type linkNode struct {
	next  *linkNode
	value string
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Length() (int, error) {
	if l.head == nil {
		return 0, errors.New("Cannot return Length, head is nil")
	}
	currNode := l.head
	c := 0
	for ; currNode != nil; c++ {
		currNode = currNode.next
	}
	return c, nil
}

func (l *LinkedList) PrintAll() ([]string, error) {
	var nodeValues []string
	if l.head == nil {
		return nodeValues, errors.New("Cannot return Length, head is nil")
	}
	currNode := l.head
	for c := 0; currNode != nil; c++ {
		nodeValues = append(nodeValues, currNode.value)
		currNode = currNode.next
	}
	return nodeValues, nil
}

func (l *LinkedList) Append(s string) error {
	node := &linkNode{nil, s}
	if l.head == nil {
		l.head = node
		return nil
	}
	currNode := l.head
	for {
		if currNode.next == nil {
			currNode.next = node
			return nil
		}
		currNode = currNode.next
	}
}

func (l *LinkedList) Swap(x int, y int) error {
	var swapX *linkNode
	var swapY *linkNode
	currNode := l.head
	for c := 0; currNode != nil; c++ {
		if x == c {
			swapX = currNode.next
		}
		if y == c {
			swapY = currNode.next
		}
		if swapX != nil && swapY != nil {
			temp := swapY
			swapY = swapX
			swapX = temp
			return nil
		}
		currNode = currNode.next
	}
	return errors.New("Index out of range")
}

func (l *LinkedList) Insert(s string, i int) error {
	if l.head == nil && i != 0 {
		return errors.New("Index out of range")
	}
	if i == 0 {
		l.head = &linkNode{l.head, s}
		return nil
	}
	var temp *linkNode
	currNode := l.head
	for c := 0; c != i; c++ {
		temp = currNode
		currNode = currNode.next
	}
	temp.next = &linkNode{currNode, s}
	return nil
}

func (l *LinkedList) Delete(i int) error {
	currNode := l.head
	if i == 0 && l.head != nil {
		l.head = l.head.next
		return nil
	}
	for c := 0; c < i-1 && currNode.next != nil; c++ {
		currNode = currNode.next
	}
	if currNode.next == nil {
		return errors.New("Index out of range")
	}
	currNode.next = currNode.next.next
	return nil
}

func (l *LinkedList) GetValueByIndex(i int) (string, error) {
	if l.head == nil {
		return "", errors.New("Index out of range")
	}
	prev := l.head
	for c := 0; c <= i; c++ {
		result := prev.value
		if c == i-1 {
			return result, nil
		}
		if prev.next == nil {
			break
		}
		prev = prev.next
	}
	return "", errors.New("Element not found")
}
func main() {
}
