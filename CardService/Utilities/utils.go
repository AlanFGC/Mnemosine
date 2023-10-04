package Utilities

import (
	"math/rand"
	"reflect"
	"time"
)

func ToInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("ToInterfaceSlice: not a slice")
	}

	result := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		result[i] = s.Index(i).Interface()
	}
	return result
}

func GenerateRandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyz"
	size := len(letters)
	randString := make([]byte, n)
	for i := 0; i < n; i++ {
		randString[i] = letters[rand.Intn(size)]
	}
	return string(randString)
}

func UniqueSlice[T any](slice1, slice2 []T) []T {
	resMap := make(map[any]bool)
	for _, item := range slice1 {
		resMap[item] = true
	}

	for _, item := range slice2 {
		resMap[item] = true
	}

	res := make([]T, 0, len(resMap))

	for item := range resMap {
		res = append(res, item.(T))
	}

	return res
}
