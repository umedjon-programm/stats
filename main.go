package main
import(
	"fmt"
)
type Money int64

type Category string

type Status string
const(
	StatusOK Status="OK"
	StatusFail Status="FAIL"
	StatusInProgress Status="INPROGRESS"
)
type Payment struct {
	ID     int
	Amount Money
	Category Category
	Status Status
}
func main(){
	fmt.Println("Hello Go")
	f:=map[Category] Money{
		"auto": 10,
		"food":20,
	}
	s:=map[Category] Money{
		"auto":10,
		"food": 25,
		"mobile":5,
		
	}
	/*cat:=[]Payment{
		{ID:1,Category:"auto",Amount: 100000},
		{ID:2,Category:"foot",Amount: 200000},
		{ID:3,Category:"auto",Amount: 300000},
		{ID:4,Category:"auto",Amount: 400000},
		{ID:5,Category:"fun",Amount: 500000},
}*/
fmt.Println(PeriodsDynamic(f,s))
}
func PeriodsDynamic(first map[Category]Money, second map[Category]Money,) map[Category]Money {
	   result:=map[Category]Money{}
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
   
func CategoriesAvg(payments []Payment)map[Category]Money{
	categoris:= map[Category]Money{}
	lencat:=map[Category]int64{}

	for _,cat:=range payments{
		categoris[cat.Category]+=cat.Amount
		lencat[cat.Category]++;
	}
	for d,c:=range categoris{
		categoris[d]=c/Money(lencat[d])
	}
	return categoris
}