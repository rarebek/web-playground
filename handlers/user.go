package handlers

import (
	"log"
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

func (h *Handlers) Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Message: "Invalid body"})
		return
	}

	if err := h.repo.InsertUser(user); err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{Message: "Internal server error happened"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "User created successfully"})
}
