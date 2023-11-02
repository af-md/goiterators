package main

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

	// what could be the things i change? iterators

	// actually change your example and add maps, filter, find
	// what are some examples that could be used?

	// could you pass a complex type?
	// what does a product type looks like?

	products := []Product{{
		name:        "perfume",
		price:       40,
		description: "feminine perfume",
	},
		{
			name:        "cigarettes",
			price:       15,
			description: "peruvian tabacco",
		},
		{
			name:        "shoes",
			price:       30,
			description: "indian kussa",
		},
	}

	fmt.Println("Product Filter iteration")
	priceRange := func(p Product) bool {
		return p.price == 30
	}
	for product := range Filter(MakeIterable(products), priceRange) {
		fmt.Println(product)
	}

	fmt.Println("Product Find iteration")
	findProduct := func(p Product) bool {
		return p.name == "perfume"
	}

	for product := range Find(MakeIterable(products), findProduct) {
		fmt.Println(product)
	}

	fmt.Println("Product Map iteration")
	mapProduct := func(p Product) bool {
		if p.name == "cigarettes" {
			p.price += 10
			return true
		}
		return false
	}
    
	for product := range Map(MakeIterable(products), mapProduct) {
		fmt.Println(product)
	}


	Try some chaining

	// fmt.Println("Product Chaining")
	// mapProductChain := func(p Product) bool {
	// 		p.price += 10
	// 		return true

	// }
	// Find(MakeIterable(products), findProduct).Map(MakeIterable(products), mapProductChain)


	// I could do potentially Wrap functions around them to do a multiple map, filter, and others sequentially
	// for item := range Map(Filter(MakeIterable(list), even), maps) {
	// fmt.Println(item)
	// }

}

// prepare the list to be iterable and take a function during the iteration

// create a function that allows the iteration to happen

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

// find returns the actual item
func Find[I any](iter Iter[I], check func(I) bool) Iter[I] {
	return func(yield func(I) bool) {
		iter(func(item I) bool  {
			if check(item) {
				yield(item)
			}
			return true
		})
	}
}

// map map doesn't return anything but instead maps the object. with what? if the integer is even then add 2 to it. its a sort of map right?

func Map[I any](iter Iter[I], check func(I) bool) Iter[I] {
	return func(yield func(I) bool) {
		iter(func(item I) bool {
			if check(item) {
				yield(item)
			}
			return true
		})
	}
}

// what's the big thing to understand here? is that you can range over a function if the function complies a certain standard!
// what's the standard here?


