package main

import (
	"./utilities"
	"log"
	"./dao/factory"
	//"./models"
	"fmt"
	//"encoding/json"
	//"time"
	"encoding/json"
	//"time"
)

func main()  {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	radicacionDao := factory.FactoryDao(config.Engine)
	radicaResult ,_ := radicacionDao.GetAll()
	//fmt.Print(radicaResult[0].NumeroRadicacion)
	fmt.Println(radicaResult)


	b, err := json.Marshal(radicaResult)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	/*parser, err := time.Parse(
		time.RFC822,
		radicaResult[0].FechaDocumento.String())

	fmt.Println("db result ",radicaResult[0].FechaDocumento.String())
	fmt.Println("parse result ",parser)*/



	/*err = userDao.Create(&user)

	if err != nil{
		log.Fatal(err)
		return
	}*/


}
