package types

type Validatable interface {
	Err() error
}
