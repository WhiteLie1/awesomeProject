package shopping

import "awesomeProject/shopping/db"
func PricceCheck(itemId int) (float64,bool) {
	item := db.LoadItem(itemId)
	if item == nil{
		return 0,false
	}
	return item.Price,true
	
}