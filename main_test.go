package main

import (
	"fmt"
	"testing"
)

// 139254	      8790 ns/op



func BenchmarkIterator(b *testing.B) {
	products := []Product{
		{
			name:        "perfume",
			price:       40,
			description: "feminine perfume",
		},
		{
			name:        "cigarettes",
			price:       15,
			description: "peruvian tobacco",
		},
		{
			name:        "shoes",
			price:       30,
			description: "indian kussa",
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		fmt.Println("Product Filter iteration")
		priceRange := func(p Product) bool {
			return p.price == 30
		}
		for product := range Filter(MakeIterable(products), priceRange) {
			fmt.Println(product)
		}
	}
}