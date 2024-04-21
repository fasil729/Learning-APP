package main

import (
    "log"
	"Brilliant/persistence/migrations"
     
    "Brilliant/leariningApp/domain"
)

func InitDB() {
    db := migrations.GetDB()

    // Create the custom user_role type
    err := db.Exec("CREATE TYPE user_role AS ENUM ('teacher', 'student', 'admin')").Error
    if err != nil {
        log.Fatalf("Error creating user_role type: %v", err)
    }

    // Check if table exists, create it if not
    if !db.Migrator().HasTable(&domain.User{}) {
        err := db.Migrator().CreateTable(&domain.User{})
        if err != nil {
            log.Fatalf("Error creating table: %v", err)
        }

        log.Println("User table created")
    } else {
        log.Println("User table already exists")
    }
}


func main() {

	InitDB()
}