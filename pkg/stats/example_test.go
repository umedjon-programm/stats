package stats

import (
	"reflect"
	"testing"

	"github.com/umedjon-programm/bank/v2/pkg/types"
)
func TestCategoriesAvg(t *testing.T) {
	
 payments:=[]types.Payment{
		 {
			 ID: 1,
			 Amount: 30000,
			 Category: "auto",
		 },
		 {
			ID: 1,
			Amount: 15000,
			Category: "food",
		},
		{
			ID: 1,
			Amount: 10000,
			Category: "auto",
		},
		{
			ID: 1,
			Amount: 20000,
			Category: "auto",
		},
		{
			ID: 1,
			Amount: 25000,
			Category: "fun",
		},

	 }
	 expected:=map[types.Category]types.Money{
		 "auto":20000,
		 "food":15000,
		 "fun":25000,
		}
		result:=CategoriesAvg(payments)
		if !reflect.DeepEqual(expected,result){
			t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
		}
	 
 }

