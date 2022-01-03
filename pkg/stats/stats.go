package stats
import(
	"github.com/umedjon-programm/bank/v2/pkg/types"
)
func Avg(payments []types.Payment) types.Money{
	var sum types.Money
	var t int
	for _, payment:= range payments{
		if payment.Status!=types.StatusFail{
		sum+=payment.Amount
		t++
		}
	}
	return sum/types.Money(t)
}
func TotalInCategory(payments []types.Payment, category types.Category) types.Money{
	var sum types.Money
	 for _,payment:=range payments{
		 if payment.Category==category {
			 if payment.Status!=types.StatusFail{
             sum+=payment.Amount
			 }
		 }
	 }
	 return sum
}