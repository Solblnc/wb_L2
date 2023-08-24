package main

type Shape interface {
	getType() string
	accept(visitor)
}
