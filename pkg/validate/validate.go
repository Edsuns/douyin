package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func Setup() {
	validate = validator.New()

	ch := zh.New()
	trans, _ = ut.New(ch, ch).GetTranslator("zh")
	err := zhs.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
}

func Struct[T any](c *gin.Context, obj *T) *T {
	var err error
	err = c.ShouldBind(obj)
	if err == nil {
		err = validate.Struct(obj)
		if err != nil {
			// if invalid
			validationErrors := err.(validator.ValidationErrors)
			var errors []string
			for _, e := range validationErrors {
				errors = append(errors, e.Translate(trans))
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, errors)
			return nil
		}
	}
	// if valid
	return obj
}
