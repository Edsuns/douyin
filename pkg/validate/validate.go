package validate

import (
	"douyin/pkg/com"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"strings"
)

var trans ut.Translator

func Setup() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// change tag name
		v.SetTagName("validate")
		// setup message translations
		registerTranslations(v)
	}
}

func registerTranslations(v *validator.Validate) {
	// use Chinese as the default translation
	ch := zh.New()
	trans, _ = ut.New(ch, ch).GetTranslator("zh")
	err := zhs.RegisterDefaultTranslations(v, trans)
	if err != nil {
		panic(err)
	}
}

func Struct[T any](c *gin.Context, obj *T) *T {
	return bind(c, c.ShouldBind, obj)
}

func StructQuery[T any](c *gin.Context, obj *T) *T {
	return bind(c, c.ShouldBindQuery, obj)
}

// bind automatically send errors when parameter errors occur
func bind[T any](c *gin.Context, bind func(obj interface{}) error, obj *T) *T {
	if err := bind(obj); err != nil {
		// if invalid
		validationErrors := err.(validator.ValidationErrors)
		var errors []string
		for _, e := range validationErrors {
			errors = append(errors, e.Translate(trans))
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, com.Response{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  strings.Join(errors, "; "),
		})
		return nil
	}
	// if valid
	return obj
}
