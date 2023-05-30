package handlers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/jinzhu/gorm"
// 	"github.com/labstack/echo/v4"
// 	"github.com/risqiikhsani/try-goecho-2/models"
// )

// type AccountHandler struct {
// 	DB *gorm.DB
// }

// func NewAccountHandler(db *gorm.DB) *AccountHandler {
// 	return &AccountHandler{
// 		DB: db,
// 	}
// }

// func (h *AccountHandler) GetUsers(c echo.Context) error {
// 	var users []models.User
// 	h.DB.Find(&users)
// 	return c.JSON(http.StatusOK, users)
// }

// func (h *AccountHandler) GetUser(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
// 	}

// 	var user models.User
// 	if err := h.DB.First(&user, id).Error; err != nil {
// 		if gorm.IsRecordNotFoundError(err) {
// 			return echo.NotFoundHandler(c)
// 		}
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// func (h *AccountHandler) CreateUser(c echo.Context) error {
// 	user := new(models.User)
// 	if err := c.Bind(user); err != nil {
// 		return err
// 	}

// 	h.DB.Create(user)

// 	return c.JSON(http.StatusCreated, user)
// }

// func (h *AccountHandler) UpdateUser(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
// 	}

// 	var user models.User
// 	if err := h.DB.First(&user, id).Error; err != nil {
// 		if gorm.IsRecordNotFoundError(err) {
// 			return echo.NotFoundHandler(c)
// 		}
// 		return err
// 	}

// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}

// 	h.DB.Save(&user)

// 	return c.JSON(http.StatusOK, user)
// }

// func (h *AccountHandler) DeleteUser(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
// 	}

// 	var user models.User
// 	if err := h.DB.First(&user, id).Error; err != nil {
// 		if gorm.IsRecordNotFoundError(err) {
// 			return echo.NotFoundHandler(c)
// 		}
// 		return err
// 	}

// 	h.DB.Delete(&user)

// 	return c.NoContent(http.StatusNoContent)
// }
