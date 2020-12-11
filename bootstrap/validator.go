package bootstrap

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/zi-ao/site-api/pkg/rule"
)

func SetupValidator() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	v.RegisterValidation("username", rule.UsernameRule)
	v.RegisterValidation("unique", rule.UniqueRule)
}
