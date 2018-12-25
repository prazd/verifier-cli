package executor

import (
	"strings"
	"fmt"
	"os"
	"github.com/prazd/go-cli/utils"
)

func Executor(userText string){
	if  userText == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}
	args := strings.Split(userText," ")
	if len(args) > 2 {
		switch args[0] {
		case "plasma":
			switch args[1] {
			case "smartContractAddress", "sca":
				if len(args) == 3 {
					fmt.Println("Smart contract address: " + args[2])
				}
			case "plasmaBalance", "pb":
				if len(args) == 3 {
					fmt.Println("Plasma balance:" + args[2])
				}
			case "smartContractBalance", "scb":
				if len(args) == 3 {
					fmt.Println("Smart contract balance:" + args[2])
				}
			case "events", "e":
				if len(args) == 3 {
					fmt.Println("Events ...")
				}
			}
		case "eth":
			switch args[1] {
			case "smartContractAddress", "sca":
				if len(args) == 3 {
					fmt.Println("Smart contract address: " + args[2])
				}
			case "smartContractBalance", "scb":
				if len(args) == 3 {
					fmt.Println("Smart contract balance:" + args[2])
				}
			case "transfer", "tr":
				if len(args) == 4{
					amountStr := args[2]
					recipientStr := args[3]
					utils.TransferETH(amountStr, recipientStr)
				}
			case "accBalance", "ab":
				if len(args) == 3{
					accStr := args[2]
					utils.AccountBalance(accStr)
				}
			}
		case "main":
			switch args[1] {
			case "deposit", "dep":
				if len(args) == 3 {
					fmt.Println("Deposit amount: " + args[2])
				}
			case "exit", "ex":
				if len(args) == 3 {
					fmt.Println("Exit func")
				}
			}
		}
	}else{
		fmt.Println("Bad args!")
	}
}
