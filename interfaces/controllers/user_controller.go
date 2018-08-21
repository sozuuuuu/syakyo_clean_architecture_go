package controllers

import (
	"strconv"

	"github.com/sozuuuuu/clean_architecture/domain"
	"github.com/sozuuuuu/clean_architecture/interfaces/database"
	"github.com/sozuuuuu/clean_architecture/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	repo := &database.UserRepository{
		SqlHandler: sqlHandler,
	}
	// なぜ&が要らない?
	interactor := usecase.UserInteractor{
		UserRepository: repo,
	}
	return &UserController{
		Interactor: interactor,
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, user)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}
