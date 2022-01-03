package stats
import(
	"github.com/umedjon-programm/bank/v2/pkg/types"
)
func CategoriesAvg(payments []types.Payment)map[types.Category]types.Money{
	categoris:= map[types.Category]types.Money{}
	lencat:=map[types.Category]int64{}

	for _,cat:=range payments{
		categoris[cat.Category]+=cat.Amount
		lencat[cat.Category]++;
	}
	for d,c:=range categoris{
		categoris[d]=c/types.Money(lencat[d])
	}
	return categoris
}