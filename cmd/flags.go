package cmd

import (
	"flag"
	"github.com/evgbay/pm/util"
)

var (
	// Flag that specifies the password length
	// Example: -l 16
	// By default: 12 symbols
	PasswordLength = flag.Int("l", util.DefaultPasswordLength, "Desired password length")
	// Flag that determines the presence of special characters when creating a password.
	// Example: -s 1 (enable)
	// By default: 0 (disable)
	EnableSpecialSymbols = flag.Int("s", util.DefaultEnableSpecialSymbolsValue, "Whether to use special characters when creating a password")
)

func ParseArgs() {
	flag.Parse()
}
