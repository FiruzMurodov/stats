package stats

import (
	"fmt"

	"github.com/FiruzMurodov/bank/v2/pkg/types"
)

func ExampleAvg()  {
	payments:= [] types.Payment{
		{
			ID: 1,
			Amount: 100,
			Category: "restaurant",
			Status: "Fail",
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
			Status: "Fail",
		},
	}

	avg:= Avg(payments)
	fmt.Println(avg)
	//Output:
	//20
}

func ExampleTotalInCategory()  {
	payments:= [] types.Payment{
		{
			ID: 1,
			Amount: 100,
			Category: "restaurant",
			Status: "Fail",
		},
		{
			ID: 2,
			Amount: 20,
			Category: "Chemist",
			Status: "Fail",
		},
		{
			ID: 3,
			Amount: 30,
			Category: "Relax",
			Status: "Fail",
		},
		{
			ID: 4,
			Amount: 70,
			Category: "Relax",
		},
		{
			ID: 5,
			Amount: 70,
			Category: "Relax",
		},
	}

	amount:= TotalInCategory(payments,"Relax")
	fmt.Println(amount)
	//Output:
	//140
}