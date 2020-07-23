package main

import (
	db "ginDemo/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8080")
}
