package main

// Function type is the generic func which gets an interface{} and returns an interface{} and an error
type Function func(v interface{}) (interface{}, error)
