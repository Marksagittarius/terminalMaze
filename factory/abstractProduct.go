package factory

import "reflect"

type AbstractProduct interface {
	GetClass() reflect.Type
	ToString() string
}
