package main

func main() {
	ferrari := GetBuilder(FerrariType)
	audi := GetBuilder(AudiType)
	factory := NewDirector(ferrari)
	fCar := factory.CreateCar()
	fCar.Print()
	factory.SetBuilder(audi)
	aCar := factory.CreateCar()
	aCar.Print()

}
