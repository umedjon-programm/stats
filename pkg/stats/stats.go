package stats
import(
	"github.com/umedjon-programm/bank/v2/pkg/types"
)
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money,) map[types.Category]types.Money {
	result:=map[types.Category]types.Money{}
	if len(first)>len(second){
	for key:=range first{
		result[key]=second[key]-first[key]
	}
 } else{
	 for key:=range second{
		 result[key]=second[key]-first[key]
	 }
 }
	return result
}