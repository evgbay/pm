package main

import (
	"fmt"
	"github.com/evgbay/pm/cmd"
	"github.com/evgbay/pm/model"
)

func main() {
	cmd.ParseArgs()
	template := model.NewTemplate(cmd.PasswordLength, cmd.EnableSpecialSymbols)
	fmt.Println(cmd.GeneratePassword(template))
}
