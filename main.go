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

}

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
