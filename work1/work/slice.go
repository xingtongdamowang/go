package main

import "errors"

func deleteSlice1[T any](slice []T, index int) ([]T, error) {
	var length = len(slice)
	if slice == nil || length == 0 {
		return nil, errors.New("slice is empty")
	}
	if index < 0 || index >= length {
		return nil, errors.New("index out of range")
	}
	for i := index; i+1 < len(slice); i++ {
		slice[i] = slice[i+1]
	}
	return slice[:length-1], nil
}

func deleteSlice2[T any](slice []T, index int) ([]T, error) {
	var length = len(slice)
	if slice == nil || length == 0 {
		return nil, errors.New("slice is empty")
	}
	if index < 0 || index >= length {
		return nil, errors.New("index out of range")
	}
	var newSlice []T
	for i := 0; i < length; i++ {
		if i == index {
			continue
		}
		newSlice = append(newSlice, slice[i])
	}
	return newSlice, nil
}
