package stats

import (
	"reflect"
	"testing"

	"github.com/umedjon-programm/bank/v2/pkg/types"
)
func TestPeriodsDynamic(t *testing.T) {
	f:=map[types.Category] types.Money{
		"auto": 10,
		"food":20,
	}
	s:=map[types.Category] types.Money{
		"auto":10,
		"food": 25,
		"mobile":5,
		
	}
	 expected:=map[types.Category]types.Money{
		 "auto":0,
		 "food":5,
		 "mobile":5,
		 
		}
		result:=PeriodsDynamic(f,s)
		if !reflect.DeepEqual(expected,result){
			t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
		}
	 
 }

