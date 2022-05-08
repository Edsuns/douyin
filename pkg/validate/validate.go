package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var trans ut.Translator

func Setup() {
	ch := zh.New()
	trans, _ = ut.New(ch, ch).GetTranslator("zh")
	// replace the default validator with our validator implementation
	binding.Validator = newValidatorImpl(trans)
}

func Struct[T any](c *gin.Context, obj *T) *T {
	return bind(c, c.ShouldBind, obj)
}

func StructQuery[T any](c *gin.Context, obj *T) *T {
	return bind(c, c.ShouldBindQuery, obj)
}

func bind[T any](c *gin.Context, bind func(obj interface{}) error, obj *T) *T {
	if err := bind(obj); err != nil {
		// if invalid
		validationErrors := err.(validator.ValidationErrors)
		var errors []string
		for _, e := range validationErrors {
			errors = append(errors, e.Translate(trans))
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errors)
		return nil
	}
	// if valid
	return obj
}
