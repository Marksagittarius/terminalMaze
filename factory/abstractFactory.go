package factory

type AbstractFactory interface {
	Produce(ProductConfig) AbstractProduct
}