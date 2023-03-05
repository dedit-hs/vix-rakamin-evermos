package migration

import (
	"fmt"
	"log"
	"mini-project/database"
	"mini-project/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database migrated.")
}