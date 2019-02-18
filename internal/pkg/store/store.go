package store

import (
)

type Value interface {}

type ModelReaderWriter interface {
	
	Store() (bool, error)
	Get() (Value, error)
	Update(Value) (error)
	Close() 
}