package admin

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"

	"../../../server/src/models"
)

func MakeAdminPanel() {
	db := getConnectionDB()
	db.AutoMigrate(
		&models.Product{}, &models.Category{}, &models.Collection{})

	// Initalize
	myAdmin := admin.New(&qor.Config{DB: db})

	// Create resources from GORM-backend model
	myAdmin.AddResource(&models.Product{})
	myAdmin.AddResource(&models.Category{})
	myAdmin.AddResource(&models.Collection{})

	// Register route
	mux := http.NewServeMux()
	// amount to /admin, so visit `/admin` to view the admin interface
	myAdmin.MountTo("/admin", mux)
	myAdmin.SetSiteName("Yime admin panel")

	log.Println("Admin panel listening on: 9000")
	http.ListenAndServe(":9000", mux)
}

func getConnectionDB() *gorm.DB {
	db, err := gorm.Open(
		"mysql", "iris:iris#max@/iris_db?charset=utf8&parseTime=True&loc=Local")
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
