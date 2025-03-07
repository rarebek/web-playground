package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rarebek/web-playground/models"
	"github.com/rarebek/web-playground/repo"
)

type Handlers struct {
	repo *repo.Repo
}

func NewHandlers(repo *repo.Repo) *Handlers {
	return &Handlers{
		repo: repo,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User registration details"
// @Success 200 {object} models.Response "User created successfully"
// @Failure 400 {object} models.Response "Invalid body"
// @Failure 500 {object} models.Response "Internal server error happened"
// @Router /register [post]
func (h *Handlers) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Message: "Invalid body"})
		return
	}
	if err := h.repo.InsertUser(user); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{Message: "Internal server error happened"})
		return
	}
	c.JSON(http.StatusOK, models.Response{Message: "User created successfully"})
}
