package cli

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/prazd/go-cli/executor"
	"github.com/prazd/go-cli/completer"
)

func CLI(){
	fmt.Println("------------Plasma Verifier----------")
	p := prompt.New(
		executor.Executor,
		completer.Completer,
		prompt.OptionPrefix("--> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}


