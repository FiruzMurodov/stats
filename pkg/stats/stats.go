package stats

import (
	"github.com/FiruzMurodov/bank/pkg/types"
)

//AVG рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	average:=0 
	for i := 0; i < len(payments); i++ {
		
		average = average + int(payments[i].Amount)
	}

	return types.Money(average/len(payments))
}

//TotalInCategory находит сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	amount:=0
	for i := 0; i < len(payments); i++ {
		if payments[i].Category==category {
			amount+=int(payments[i].Amount)
		}
	}
	return types.Money(amount)
}