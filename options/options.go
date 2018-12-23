package options

import "github.com/c-bata/go-prompt"


var MainOptions = []prompt.Suggest{
	prompt.Suggest{Text: "get", Description: "Get option"},
	prompt.Suggest{Text: "exit", Description: "Exit from app"},

}


var GetOptions = []prompt.Suggest{
	prompt.Suggest{Text: "smartContractAddress", Description: "sca: Get Smart Contract address."},
	prompt.Suggest{Text: "plasmaBalance", Description: "pb: Get balance of my account in Plasma."},
	prompt.Suggest{Text: "smartContractBalance", Description: "scb: Get balance of Plasma smart contract."},
	prompt.Suggest{Text: "events", Description: "e: Get all events.Reduction"},
}
