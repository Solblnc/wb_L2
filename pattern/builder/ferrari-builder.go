package main

type Ferrari struct {
	Brand        string
	Engine       int
	Brakes       string
	Transmission string
}

func (f *Ferrari) SetBrand() {
	f.Brand = "ferrari"
}

func (f *Ferrari) SetEngine() {
	f.Engine = 720
}

func (f *Ferrari) SetBrakes() {
	f.Brakes = "brambo"
}

func (f *Ferrari) SetTransmission() {
	f.Transmission = "2-robot"
}

func (f *Ferrari) GetCar() Car {
	return Car{
		Brand:        f.Brand,
		Engine:       f.Engine,
		Brakes:       f.Brakes,
		Transmission: f.Transmission,
	}

}
