package cmd

import (
	"github.com/evgbay/pm/model"
	"math/rand"
	"strings"
)

func GeneratePassword(template *model.Template) string {
	chars := model.GetCharset(template)
	var builder strings.Builder
	builder.Grow(template.Len)
	for i := 0; i < template.Len; i++ {
		builder.WriteRune(chars[rand.Intn(len(chars))])
	}
	return builder.String()
}
