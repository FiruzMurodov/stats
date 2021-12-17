package stats

import (
	"fmt"

	"github.com/FiruzMurodov/bank/pkg/types"
)

func ExampleAvg()  {
	payments:= [] types.Payment{
		{
			ID: 1,
			Amount: 100,
			Category: "restaurant",
		},
		{
			ID: 2,
			Amount: 20,
			Category: "Chemist",
		},
		{
			ID: 3,
			Amount: 30,
			Category: "Relax",
		},
	}

	avg:= Avg(payments)
	fmt.Println(avg)
	//Output:
	//50
}

func ExampleTotalInCategory()  {
	payments:= [] types.Payment{
		{
			ID: 1,
			Amount: 100,
			Category: "restaurant",
		},
		{
			ID: 2,
			Amount: 20,
			Category: "Chemist",
		},
		{
			ID: 3,
			Amount: 30,
			Category: "Relax",
		},
		{
			ID: 4,
			Amount: 70,
			Category: "Relax",
		},
	}

	amount:= TotalInCategory(payments,"Relax")
	fmt.Println(amount)
	//Output:
	//100
}