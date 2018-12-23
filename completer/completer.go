package completer

import (
	"github.com/c-bata/go-prompt"
	"strings"
	"../options"
)

func Completer(d prompt.Document) []prompt.Suggest{
	args := strings.Split(d.TextBeforeCursor(), " ")
	return argumentsCompleter(args)
}

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(options.MainOptions, args[0], true)
	}
	first := args[0]
	switch first{
	case "get":
		second := args[1]
		if len(args) == 2{
			return prompt.FilterHasPrefix(options.GetOptions, second, true)
		}
	}
return []prompt.Suggest{}
}