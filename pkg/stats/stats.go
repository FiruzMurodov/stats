package stats

import (
	"github.com/FiruzMurodov/bank/v2/pkg/types"
)

//AVG рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {

	average := 0
	avg := 0
	for i := 0; i < len(payments); i++ {
		if payments[i].Status != types.StatusFail {
			average += int(payments[i].Amount)
			avg += 1
		}
	}

	return types.Money(average / avg)
}

//TotalInCategory находит сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	amount := 0
	for i := 0; i < len(payments); i++ {
		if payments[i].Category == category && payments[i].Status != types.StatusFail {
			amount += int(payments[i].Amount)
		}
	}
	return types.Money(amount)
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {

	avg := make(map[types.Category]types.Money)
	index := make(map[types.Category]types.Money)

	for _, payment := range payments {
		_, key := avg[payment.Category]

		if key {
			avg[payment.Category] += payment.Amount
			index[payment.Category]++
		} else {
			avg[payment.Category] = payment.Amount
			index[payment.Category] = 1
		}

	}
	for key, value := range avg {
		avg[key] = value / index[key]
	}

	return avg
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {

	result := make(map[types.Category]types.Money)

	if len(first) > len(second) {

		for key := range first {
			result[key] = second[key] - first[key]
		}

		return result
	}

	for key := range second {
		result[key] = second[key] - first[key]
	}

	return result

}
