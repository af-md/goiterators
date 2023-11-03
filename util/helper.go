package helper

import (
	"fmt"
)

type Iter[I any] func(func(yield I) bool)

type Product struct {
	name        string
	price       int
	description string
}

func main() {

	products := []Product{{
		name:        "perfume",
		price:       40,
		description: "paco rabanne collection",
	},
		{
			name:        "cigarettes",
			price:       15,
			description: "peruvian tabacco",
		},
		{
			name:        "shoes",
			price:       30,
			description: "comfortable trainers",
		},
	}

	fmt.Println("Product Filter iteration")
	priceRange := func(p Product) bool {
		return p.price == 30
	}
	for product := range Filter(MakeIterable(products), priceRange) {
		fmt.Println(product)
	}

	// fmt.Println("Product Find iteration")
	// findProduct := func(p Product) bool {
	// 	return p.name == "perfume"
	// }

	// for product := range Find(MakeIterable(products), findProduct); ok {
	// 		fmt.Println(product)
	// }

	fmt.Println("Product Map iteration")
	mapProduct := func(p Product) Product {
		if p.name == "perfume" {
			p.price = p.price + 10
			return p
		}
		return p
	}

	for product := range Map(MakeIterable(products), mapProduct) {
		fmt.Println(product)
	}
}

// the item, []. in your code you'll check if item is empty? or check a boolean value and the item?
// func Find[I any](iter Iter[I], match func(I) bool) (result I, ok bool) {
// 	for product := range iter {
// 		if match(product) {
// 			return product, true
// 		}
// 	}
// 	return result, false

// }

func MakeIterable[S ~[]I, I any](s S) Iter[I] {
	return func(yield func(I) bool) {
		for _, item := range s {
			yield(item)
		}
	}
}

func Filter[I any](iter Iter[I], check func(I) bool) Iter[I] {
	return func(yield func(I) bool) {
		iter(func(item I) bool {
			if check(item) {
				yield(item)
			}
			return true
		})
	}
}

func Map[I any, T any](iter Iter[I], mapf func(I) T) Iter[T] {
	return func(yield func(T) bool) {
		iter(func(item I) bool {
			return yield(mapf(item))
		})
	}
}

func HELP() bool  {
	return true
}
