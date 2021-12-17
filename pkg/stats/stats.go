package stats

import (
	"github.com/FiruzMurodov/bank/v2/pkg/types"
)

//AVG рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {

	average:=0 
	avg:=0
	for i := 0; i < len(payments); i++ {
		if payments[i].Status != types.StatusFail {
			average += int(payments[i].Amount)
			avg+=1
		}
	}

	return types.Money(average/avg)
}

//TotalInCategory находит сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	amount:=0
	for i := 0; i < len(payments); i++ {
		if payments[i].Category==category && payments[i].Status!=types.StatusFail{
			amount+=int(payments[i].Amount)
		}
	}
	return types.Money(amount)
}