package controllers

import (
	"github.com/gin-gonic/gin"
	"go-base/pkg/crud/core/entities"
	"go-base/pkg/crud/core/usecases"
	"go-base/pkg/crud/restful/requests"
	"go-base/pkg/helpers/gin_custom"
	"go-base/pkg/helpers/log"
	"net/http"
)

type ExampleController struct {
	*gin_custom.BaseController
	exampleUseCase *usecases.ExampleUseCase
}

func NewExampleController(baseController *gin_custom.BaseController, exampleUseCase *usecases.ExampleUseCase) *ExampleController {
	return &ExampleController{baseController, exampleUseCase}
}

func (ctrl *ExampleController) Create(c *gin.Context) {
	var req requests.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf(err, "bind json fail")
		ctrl.DefaultError(c, http.StatusBadRequest)
		return
	}

	if err := ctrl.ValidateRequest(req); err != nil {
		ctrl.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := ctrl.exampleUseCase.Create(&entities.Example{Name: req.Name}); err != nil {
		log.Errorf(err, "Create repositories fail")
		ctrl.DefaultError(c, http.StatusBadRequest)
		return
	}
	ctrl.Success(c, "")
}

func (ctrl *ExampleController) Read(c *gin.Context) {
	listExample, err := ctrl.exampleUseCase.Read()
	if err != nil {
		log.Errorf(err, "Read repositories fail")
		ctrl.DefaultError(c, http.StatusBadRequest)
		return
	}
	ctrl.Success(c, listExample)
}

func (ctrl *ExampleController) Update(c *gin.Context) {
	var req requests.EditRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf(err, "bind json fail")
		ctrl.DefaultError(c, http.StatusBadRequest)
		return
	}
	if err := ctrl.ValidateRequest(req); err != nil {
		ctrl.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := ctrl.exampleUseCase.Update(req.Id, &entities.Example{Name: req.Name}); err != nil {
		log.Errorf(err, "Update repositories fail")
		ctrl.DefaultError(c, http.StatusBadRequest)
		return
	}
	ctrl.Success(c, "")
}

func (ctrl *ExampleController) Delete(c *gin.Context) {
	var req requests.DeleteRequest
	if err := ctrl.exampleUseCase.Delete(req.Id); err != nil {
		log.Errorf(err, "Delete repositories fail")
		ctrl.DefaultError(c, http.StatusBadRequest)
		return
	}
	ctrl.Success(c, "")
}
