package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	 var mapPerson  = make(map[string]string)
	mapPerson["user_1"] ="address_1"
	mapPerson["user_2"] = "address_2"
	//fmt.Println(mapPerson)
	zArray,error := json.Marshal(mapPerson)
	if(error == nil){
		fmt.Println(string(zArray))
	}else{
		fmt.Println("Error in Json marshaling")
	}

}