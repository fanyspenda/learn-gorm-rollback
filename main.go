package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OnFailError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err.Error())
	}
}

func initDB() *gorm.DB {
	dsn := "host= user= password= dbname= port= sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	OnFailError(err, "Error Connecting to DB")
	return db
}

func main() {
	db := initDB()
	db.Transaction(
		func(tx *gorm.DB) error {

			for i := 0; i < 2; i++ {

				query := `insert into "user" (email, "name", address) 
				values ('johny` + fmt.Sprintf("%d", i) + `@gmail.com', 'Johny', 'Jember');`

				if err := tx.Exec(query, i).Error; err != nil {
					OnFailError(err, "error inserting data")
					return err
				}
			}
			tx.Rollback()
			return nil
		},
	)
}
