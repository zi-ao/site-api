package rule

import (
	"github.com/go-playground/validator/v10"
)

// illegalUsernames 用户名黑名单
var illegalUsernames = []string{
	"admin",
	"admins",
	"root",
	"test",
	"administrator",
	"administrators",
}

// UsernameRule 用户名规则
func UsernameRule(field validator.FieldLevel) bool {
	username := field.Field().String()
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	start, end := 0, len(illegalUsernames)-1
	for start < end {
		if illegalUsernames[start] == username || illegalUsernames[end] == username {
			return false
		}
		start++
		end--
	}
	if start == end && illegalUsernames[start] == username {
		return false
	}
	return true
}
