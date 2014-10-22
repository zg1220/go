package loginController

import (
	"../models"
	"log"
	"menteslibres.net/gosexy/redis"
)

var client *redis.Client
var err error
var host = "127.0.0.1"
var port uint = 6379

func Connect() {
	client = redis.New()
	err = client.Connect(host, port)
	if err != nil {
		log.Fatalf("Connect failed: %s\n", err.Error())
		return
	}
}
func Select(name string) string {
	Connect()
	pwd, err := client.Get(name)
	if err != nil {
		pwd = "3"
	}
	client.Quit()
	return pwd
}

func Insert(u loginModel.Users) {
	print(u.Name)
	// Connect()
	// i, err := client.Set(u.Name, u.Pwd)
	// if err != nil {
	// 	print(err)
	// }
	// client.Quit()
	// return i
}
