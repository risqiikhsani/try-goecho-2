package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/risqiikhsani/try-goecho-2/database"
	"github.com/risqiikhsani/try-goecho-2/handlers"
)

var (
	db *gorm.DB
)

func main() {

	// initialize database
	db = database.ConnectDB()
	defer db.Close()

	// initialize the echo server
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// use handlers
	userHandler := handlers.NewUserHandler(db)

	// routes
	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)

	// log
	log.Println("Server started on http://localhost:8080")

	// start the server
	e.Start(":8080")
}

// func initDB() *gorm.DB {
// 	db, err := gorm.Open("mysql", "username:password@tcp(127.0.0.1:3306)/percobaan?charset=utf8mb4&parseTime=True&loc=Local")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// db.AutoMigrate(&User{})
// 	// use models
// 	db.AutoMigrate(&models.User{})

// 	return db
// }

// func GetUsers(c echo.Context) error {
// 	var users []User
// 	db.Find(&users)
// 	return c.JSON(http.StatusOK, users)
// }

// func GetUser(c echo.Context) error {
// 	id := c.Param("id")

// 	var user User
// 	if err := db.First(&user, id).Error; err != nil {
// 		if gorm.IsRecordNotFoundError(err) {
// 			return echo.NotFoundHandler(c)
// 		}
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// func CreateUser(c echo.Context) error {
// 	user := new(User)
// 	if err := c.Bind(user); err != nil {
// 		return err
// 	}

// 	db.Create(user)

// 	return c.JSON(http.StatusCreated, user)
// }

// func UpdateUser(c echo.Context) error {
// 	id := c.Param("id")

// 	var user User
// 	if err := db.First(&user, id).Error; err != nil {
// 		if gorm.IsRecordNotFoundError(err) {
// 			return echo.NotFoundHandler(c)
// 		}
// 		return err
// 	}

// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}

// 	db.Save(&user)

// 	return c.JSON(http.StatusOK, user)
// }

// func DeleteUser(c echo.Context) error {
// 	id := c.Param("id")

// 	var user User
// 	if err := db.First(&user, id).Error; err != nil {
// 		if gorm.IsRecordNotFoundError(err) {
// 			return echo.NotFoundHandler(c)
// 		}
// 		return err
// 	}

// 	db.Delete(&user)

// 	return c.NoContent(http.StatusNoContent)
// }
