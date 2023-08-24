package main

type Audi struct {
	Brand        string
	Engine       int
	Brakes       string
	Transmission string
}

func (a *Audi) SetBrand() {
	a.Brand = "audi"
}

func (a *Audi) SetEngine() {
	a.Engine = 450
}

func (a *Audi) SetBrakes() {
	a.Brakes = "audi-sport"
}

func (a *Audi) SetTransmission() {
	a.Transmission = "ZF"
}

func (a *Audi) GetCar() Car {
	return Car{
		Brand:        a.Brand,
		Engine:       a.Engine,
		Brakes:       a.Brakes,
		Transmission: a.Transmission,
	}

}
