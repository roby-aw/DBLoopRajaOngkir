package main

import (
	"dblooprajaongkir/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func main() {
	dsn := "root123:root123@tcp(localhost:3306)/dbcity?parseTime=true"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// DB.AutoMigrate(&models.Province{})

	for i := 2; i <= 34; i++ {
		url := fmt.Sprintf("https://api.rajaongkir.com/starter/province?id=%d", i)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("key", "75ac21477faeddb1b34f8dd63fc61f45")
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		hasil := string(body)
		var city models.CostResponse

		err = json.Unmarshal([]byte(hasil), &city)
		if err != nil {
			fmt.Println(err)
		}
		arr := city.Rajaongkir.Results
		arr.ID = i
		fmt.Println(arr)
		DB.Create(arr)
	}
}

func dbmigrate() {
	DB.AutoMigrate(&models.City{})
}

// func city
// func main() {
// 	dsn := "root123:root123@tcp(localhost:3306)/dbcity?parseTime=true"
// 	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	DB.AutoMigrate(&models.Results{})

// 	for i := 393; i <= 501; i++ {
// 		url := fmt.Sprintf("https://api.rajaongkir.com/starter/city?id=%d", i)
// 		req, _ := http.NewRequest("GET", url, nil)
// 		req.Header.Add("key", "75ac21477faeddb1b34f8dd63fc61f45")
// 		res, _ := http.DefaultClient.Do(req)
// 		defer res.Body.Close()
// 		body, _ := ioutil.ReadAll(res.Body)
// 		hasil := string(body)
// 		var city models.CostResponse

// 		err = json.Unmarshal([]byte(hasil), &city)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		arr := city.Rajaongkir.Results
// 		arr.ID = i
// 		fmt.Println(arr)
// 		DB.Create(arr)
// 	}
// }
