// package main

// import "strings"

// type Iter[I any] func(yield func(I) bool) bool

// func main()  {

// 	list := [4]int{1,2,3,4}

// }

// func MakeIterable[S ~[]I, I any](s S) Iter[I] {
// 	return func(yield func(I) bool) {
// 		for _, item := range s {
// 			yield(item)
// 		}
// 	}
// }