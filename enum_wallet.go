package betting

var EWallet = eWallet{
	UK:  "UK",
	AUS: "AUS",
}

type eWallet struct {
	UK  eWalletInternal
	AUS eWalletInternal
}

type eWalletInternal string
