package main

type Director struct {
	Builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{Builder: builder}
}

func (f *Director) SetBuilder(builder Builder) {
	f.Builder = builder
}

func (f *Director) CreateCar() Car {
	f.Builder.SetBrand()
	f.Builder.SetEngine()
	f.Builder.SetBrakes()
	f.Builder.SetTransmission()
	return f.Builder.GetCar()
}
