package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {

	result := card.Withdraw(types.Card{Balance: 10_000_00, Active: true}, 500_00)
	fmt.Println(result.Balance)
}
