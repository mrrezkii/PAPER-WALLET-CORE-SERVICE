package shared

import v10 "github.com/go-playground/validator/v10"

type (
	Validator interface {
		Validate(interface{}) error
	}

	Validate struct {
		instance *v10.Validate
	}
)

func NewValidator() *Validate {
	v := v10.New(v10.WithRequiredStructEnabled()) // Tambahkan opsi yang diinginkan
	return &Validate{instance: v}
}

func (v *Validate) Validate(i interface{}) error {
	return v.instance.Struct(i)
}
