package main

const (
	FerrariType = "ferrari"
	AudiType    = "audi"
)

type Builder interface {
	SetBrand()
	SetEngine()
	SetBrakes()
	SetTransmission()
	GetCar() Car
}

func GetBuilder(car string) Builder {
	switch car {
	default:
		return nil
	case FerrariType:
		return &Ferrari{}
	case AudiType:
		return &Audi{}
	}
}
