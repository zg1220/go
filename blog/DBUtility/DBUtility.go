package DBUtility

import (
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
		log.Fatalf("Could not GET: %s\n", err.Error())
	}
	client.Quit()
	return pwd
}
