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
	
	func TestPeriodsDynamic(t *testing.T) {
		first := map[types.Category]types.Money {
				"auto": 10,
			    "food":20,
			
		}

		second:= map[types.Category]types.Money {
			"auto": 10,
			"food": 25,
			"mobile":5,
		
	}
		
		excepted := map[types.Category]types.Money {
			"auto":0,
			"food":5,
			"mobile":5,
			
		}

		result:= PeriodsDynamic(first,second)

		if !reflect.DeepEqual(excepted,result) {
			t.Errorf("invalid results matching map excepted %v,actual %v",excepted, result)
		}
	
	}