package stats

import (
	"reflect"
	"testing"

	"github.com/FiruzMurodov/bank/v2/pkg/types")

	func TestCategoriesAvg(t *testing.T) {
		payments :=  []types.Payment{
			{
				ID:  1,
				Amount:  10,
				Category: "auto",
			},
			{
				ID :  2,
				Amount : 20,
				Category: "auto",
			},
			{
				ID:  3,
				Amount: 10,
				Category: "food",
			},
				{
				ID:  4,
				Amount: 30,
				Category: "food",
			},
				{
				ID:  4,
				Amount: 10,
				Category: "fun",
			},
	
		}
		excepted := map[types.Category]types.Money {
			"auto":15,
			"food":20,
			"fun":10,
		}

		result:= CategoriesAvg(payments)

		if !reflect.DeepEqual(excepted,result) {
			t.Errorf("invalid results matching map excepted %v,actual %v",excepted, result)
		}
	
	}