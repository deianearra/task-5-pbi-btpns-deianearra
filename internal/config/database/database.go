// package database

// import (
//     "gorm.io/driver/mysql" // Sesuaikan dengan database yang digunakan
//     "gorm.io/gorm"
// )

//	func ConnectDB() (*gorm.DB, error) {
//	    dsn := "username:password@tcp(localhost:3306)/nama_database?charset=utf8mb4&parseTime=True&loc=Local" // Sesuaikan dengan konfigurasi database yang digunakan
//	    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	    if err != nil {
//	        return nil, err
//	    }
//	    return db, nil
//	}
// package database

// import (
// 	"task-5-pbi-btpns-deianearra/internal/models"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Connect() {
// 	dsn := "your_database_connection_string" // Add your database connection string here

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to connect to database")
// 	}

// 	DB = db

// 	// Auto-migrate the models
// 	err = DB.AutoMigrate(&models.User{}, &models.Photo{})
// 	if err != nil {
// 		panic("Failed to perform auto migration")
// 	}
// }

package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"task-5-pbi-btpns-deianearra/internal/models"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
		return nil, err
	}

	DB = db
	err = DB.AutoMigrate(&models.User{}, &models.Photo{})
	if err != nil {
		panic("Failed to perform auto migration")
	}
	return db, nil
}

func Close() {
	db, _ := DB.DB()
	db.Close()
}
