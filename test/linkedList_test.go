package tests

import (
	"Escargot/structures/lists"
	"fmt"
	"math/rand"
	"testing"
)

func TestEmptyList(t *testing.T) {
	list := lists.LinkedList{}
	if node := list.GetHead(); node != nil {
		t.Errorf("Got a non-nil head when empty. Node is %v", node)
	}
	if node := list.GetTail(); node != nil {
		t.Errorf("Got a non-nil head when empty. Node is %v", node)
	}
	if slice := list.Slice(); len(slice) != 0 {
		t.Errorf("LinkedList not empty. Contents are %v", slice)
	}
}

func TestOneItemListAppend(t *testing.T) {
	list := lists.LinkedList{}
	data := 5
	list.Append(fmt.Sprintf("%v", data), data)
	if nodeData := list.GetHead(); nodeData == nil {
		t.Errorf("Got nil as the head of a LinkedList with one item in it.")
	} else if nodeData != data {
		t.Errorf("Got an unexcepted value in the head of a single item linked list. Expected %v  but got %v", data, nodeData)
	}
	if nodeData := list.GetTail(); nodeData == nil {
		t.Errorf("Got nil as the tail of a LinkedList with one item in it.")
	} else if nodeData != data {
		t.Errorf("Got an unexcepted value in the tail of a single item linked list. Expected %v  but got %v", data, nodeData)
	}
	if slice := list.Slice(); len(slice) != 1 {
		t.Errorf("LinkedList does not have exactly one item in it. Contents are %v", slice)
	}
}

func TestOneItemListPush(t *testing.T) {
	list := lists.LinkedList{}
	data := 5
	list.Push(fmt.Sprintf("%v", data), data)
	if nodeData := list.GetHead(); nodeData == nil {
		t.Errorf("Got nil as the head of a LinkedList with one item in it.")
	} else if nodeData != data {
		t.Errorf("Got an unexcepted value in the head of a single item linked list. Expected %v  but got %v", data, nodeData)
	}
	if nodeData := list.GetTail(); nodeData == nil {
		t.Errorf("Got nil as the tail of a LinkedList with one item in it.")
	} else if nodeData != data {
		t.Errorf("Got an unexcepted value in the tail of a single item linked list. Expected %v  but got %v", data, nodeData)
	}
	if slice := list.Slice(); len(slice) != 1 {
		t.Errorf("LinkedList does not have exactly one item in it. Contents are %v", slice)
	}
}

func TestTwoItemListAppend(t *testing.T) {
	list := lists.LinkedList{}
	dataOne := 5
	dataTwo := "10"
	list.Append(fmt.Sprintf("%v", dataOne), dataOne)
	list.Append(fmt.Sprintf("%v", dataTwo), dataTwo)
	if nodeData := list.GetHead(); nodeData == nil {
		t.Errorf("Got nil as the head of a LinkedList with two items in it.")
	} else if nodeData != dataOne {
		t.Errorf("Got an unexcepted value in the head of a single item linked list. Expected %v but got %v", dataOne, nodeData)
	}
	if nodeData := list.GetTail(); nodeData == nil {
		t.Errorf("Got nil as the tail of a LinkedList with one item in it.")
	} else if nodeData != dataTwo {
		t.Errorf("Got an unexcepted value in the tail of a single item linked list. Expected %v but got %v", dataTwo, nodeData)
	}
	targetSlice := []interface{}{dataOne, dataTwo}
	resultSlice := list.Slice()
	if len(resultSlice) != len(targetSlice) {
		t.Errorf("LinkedList is the incorrect length. Expected %v, got %v", len(targetSlice), len(resultSlice))
	}
	for index, value := range list.Slice() {
		if value != targetSlice[index] {
			t.Errorf("Unexpected value and index %v. Expected %v, got %v", index, targetSlice[index], value)
		}
	}
}

func TestTwoItemListPush(t *testing.T) {
	list := lists.LinkedList{}
	dataOne := 5
	dataTwo := "10"
	list.Push(fmt.Sprintf("%v", dataOne), dataOne)
	list.Push(fmt.Sprintf("%v", dataTwo), dataTwo)
	if nodeData := list.GetHead(); nodeData == nil {
		t.Errorf("Got nil as the head of a LinkedList with two items in it.")
	} else if nodeData != dataTwo {
		t.Errorf("Got an unexcepted value in the head of a single item linked list. Expected %v but got %v", dataOne, nodeData)
	}
	if nodeData := list.GetTail(); nodeData == nil {
		t.Errorf("Got nil as the tail of a LinkedList with one item in it.")
	} else if nodeData != dataOne {
		t.Errorf("Got an unexcepted value in the tail of a single item linked list. Expected %v but got %v", dataTwo, nodeData)
	}
	targetSlice := []interface{}{dataTwo, dataOne}
	resultSlice := list.Slice()
	if len(resultSlice) != len(targetSlice) {
		t.Errorf("LinkedList is the incorrect length. Expected %v, got %v", len(targetSlice), len(resultSlice))
	}
	for index, value := range list.Slice() {
		if value != targetSlice[index] {
			t.Errorf("Unexpected value and index %v. Expected %v, got %v", index, targetSlice[index], value)
		}
	}
}

func TestPopEmpty(t *testing.T) {
	result := (&lists.LinkedList{}).Pop()
	if result != nil {
		t.Errorf("Popped empty list. Expected nil, got %v", result)
	}
}

func TestPopSingleItemAppend(t *testing.T) {
	list := lists.LinkedList{}
	data := 5
	list.Append(fmt.Sprintf("%v", data), data)
	if result := list.Pop(); result != data {
		t.Errorf("Incorrect pop. Expected %v, got %v", data, result)
	}
	if result := list.Pop(); result != nil {
		t.Errorf("Incorrect pop. Expected %v, got %v", nil, result)
	}
}

func TestPopSingleItemPush(t *testing.T) {
	list := lists.LinkedList{}
	data := 5
	list.Push(fmt.Sprintf("%v", data), data)
	if result := list.Pop(); result != data {
		t.Errorf("Incorrect pop. Expected %v, got %v", data, result)
	}
	if result := list.Pop(); result != nil {
		t.Errorf("Incorrect pop. Expected %v, got %v", nil, result)
	}
}

func TestPopMultipleItemsAppend(t *testing.T) {
	list := lists.LinkedList{}
	length := 1000
	var numbers []int
	for i := 0; i < length; i++ {
		random := rand.Int()
		numbers = append(numbers, random)
		list.Append(fmt.Sprintf("%v", random), random)
	}
	for _, value := range numbers {
		if result := list.Pop(); result != value {
			t.Errorf("Incorrect pop. Expected %v, got %v", value, result)
		}
	}
	if result := list.Pop(); result != nil {
		t.Errorf("Incorrect pop. Expected %v, got %v", nil, result)
	}
}

func TestPopMultipleItemsPush(t *testing.T) {
	list := lists.LinkedList{}
	length := 1000
	var numbers []int
	for i := 0; i < length; i++ {
		random := rand.Int()
		numbers = append(numbers, random)
		list.Push(fmt.Sprintf("%v", random), random)
	}
	for index := len(numbers) - 1; index >= 0; index-- {
		value := numbers[index]
		if result := list.Pop(); result != value {
			t.Errorf("Incorrect pop. Expected %v, got %v", value, result)
		}
	}
	if result := list.Pop(); result != nil {
		t.Errorf("Incorrect pop. Expected %v, got %v", nil, result)
	}
}

func TestPopMultipleItemsAppendPush(t *testing.T) {
	list := lists.LinkedList{}
	length := 1000
	var numbers []int
	for i := 0; i < length/2; i++ {
		random := rand.Int()
		numbers = append(numbers, random)
		list.Append(fmt.Sprintf("%v", random), random)
	}
	for i := 0; i < length/2; i++ {
		random := rand.Int()
		numbers = append(numbers, random)
		list.Push(fmt.Sprintf("%v", random), random)
	}
	for i := len(numbers) - 1; i >= len(numbers)/2; i-- {
		value := numbers[i]
		if result := list.Pop(); result != value {
			t.Errorf("Incorrect pop. Expected %v, got %v", value, result)
		}
	}
	for i := 0; i < len(numbers)/2; i++ {
		value := numbers[i]
		if result := list.Pop(); result != value {
			t.Errorf("Incorrect pop. Expected %v, got %v", value, result)
		}
	}
	if result := list.Pop(); result != nil {
		t.Errorf("Incorrect pop. Expected %v, got %v", nil, result)
	}
}
