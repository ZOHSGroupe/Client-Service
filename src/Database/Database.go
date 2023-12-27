package Database

import (
	"AUTH-SERVICE/src/Models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitialiserBaseDeDonnees() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Impossible de se connecter à la base de données")
	}
	DB = db
	fmt.Println("Connecté à la base de données")
	//err = db.Table("client").AutoMigrate(&Models.Client{})
	//if err != nil {
	//panic("Erreur lors de la migration de la table 'client'")
	//}
	db.AutoMigrate(&Models.Client{})

}
