package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000

}
func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 0

}
func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000

}
func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 1000_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000

}
func ExampleDeposit_normal() {

	result := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&result, 10_000_00)
	fmt.Println(result.Balance)
	//Output: 3000000

}
func ExampleDeposit_inactive() {
	result := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&result, 10_000_00)
	fmt.Println(result.Balance)
	//Output: 2000000

}
func ExampleDeposit_limit() {
	result := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&result, 60_000_00)
	fmt.Println(result.Balance)
	//Output: 2000000
}

func ExampleAddBonus_normal() {
	result := types.Card{Balance: 20_000_00, MinBalance: 1_000_00, Active: true}
	AddBonus(&result, 3, 30, 365)
	fmt.Println(result.Balance)
	//Output: 2000246
}

func ExampleAddBonus_inactive() {
	result := types.Card{Balance: 20_000_00, MinBalance: 1_000_00, Active: false}
	AddBonus(&result, 3, 30, 365)
	fmt.Println(result.Balance)
	//Output: 2000000
}

func ExampleAddBonus_negative() {
	result := types.Card{Balance: 0, MinBalance: 1_000_00, Active: true}
	AddBonus(&result, 3, 30, 365)
	fmt.Println(result.Balance)
	//Output: 0
}

func ExampleAddBonus_limit() {
	result := types.Card{Balance: 1, MinBalance: 1202_777_877, Active: true}
	AddBonus(&result, 3, 30, 365)
	fmt.Println(result.Balance)
	//Output: 1
}
func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 20_000_00,
			Active:  true,
		},
		{
			Balance: -20_000_00,
			Active:  true,
		},
		{
			Balance: 20_000_00,
			Active:  false,
		},
	}
	fmt.Println(Total(cards))
	//Output: 2000000
}
func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Active:  true,
			PAN:     "5584",
			Balance: types.Money(1_000_00),
		},
		{
			Active:  false,
			PAN:     "1844",
			Balance: types.Money(1_000_00),
		},
		{
			Active:  true,
			PAN:     "8488",
			Balance: types.Money(-1_000_00),
		},
	}

	sources := PaymentSources(cards)

	for _, source := range sources {
		fmt.Println(source.Number)
	}
	// Output:
	// 5584
}