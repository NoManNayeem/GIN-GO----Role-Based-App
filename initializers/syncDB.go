package initializers

import "role_based_app/models"

func SyncDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate the database: " + err.Error())
	}
}
