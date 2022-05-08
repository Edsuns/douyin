package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
)

type validatorImpl struct {
	validate *validator.Validate
}

func newValidatorImpl(trans ut.Translator) *validatorImpl {
	v := &validatorImpl{}

	v.validate = validator.New()

	err := zhs.RegisterDefaultTranslations(v.validate, trans)
	if err != nil {
		panic(err)
	}

	return v
}

// ValidateStruct validates a struct
func (v *validatorImpl) ValidateStruct(obj interface{}) error {
	return v.validate.Struct(obj)
}

func (v *validatorImpl) Engine() interface{} {
	return v.validate
}
