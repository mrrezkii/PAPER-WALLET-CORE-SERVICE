package shared

import v10 "github.com/go-playground/validator/v10"

type (
	Validator interface {
		Validate(interface{}) error
	}

	validate struct {
		instance *v10.Validate
	}
)

func NewValidator() *validate {
	v := v10.New(v10.WithRequiredStructEnabled()) // Tambahkan opsi yang diinginkan
	return &validate{instance: v}
}

func (v *validate) Validate(i interface{}) error {
	return v.instance.Struct(i)
}
