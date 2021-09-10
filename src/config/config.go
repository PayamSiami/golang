package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	PORT     = 0
	DBURL    = ""
	DBDRIVER = ""
)

func Load() {
	var err error
	// err = godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 9000
	}

	DBDRIVER = "mysql"
	DBURL = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"foo", "p@ssW0rd", "blogos")
}
