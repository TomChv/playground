package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct {
	repo *Repository
}

func NewController(repo *Repository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) Create(g *gin.Context) {
	var data CreateData

	if err := g.ShouldBind(&data); err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	newTask, err := c.repo.Create(g, &data)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	g.JSON(http.StatusCreated, newTask)
}

func (c *Controller) List(g *gin.Context) {
	tasks, err := c.repo.List(g)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	g.JSON(http.StatusOK, tasks)
}

func (c *Controller) Get(g *gin.Context) {
	id, err := uuid.Parse(g.Param("id"))
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
	}

	task, err := c.repo.Get(g, id)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	g.JSON(http.StatusOK, task)
}

func (c *Controller) Update(g *gin.Context) {
	id, err := uuid.Parse(g.Param("id"))
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
	}

	var data UpdateData
	if err := g.ShouldBind(&data); err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	task, err := c.repo.Update(g, id, &data)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	g.JSON(http.StatusOK, task)
}

func (c *Controller) Delete(g *gin.Context) {
	id, err := uuid.Parse(g.Param("id"))
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
	}

	err = c.repo.Delete(g, id)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	g.String(http.StatusOK, "task successfully deleted")
}

func (c *Controller) Start(g *gin.Context) {
	id, err := uuid.Parse(g.Param("id"))
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
	}

	task, err := c.repo.Start(g, id)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	g.JSON(http.StatusOK, task)
}

func (c *Controller) Done(g *gin.Context) {
	id, err := uuid.Parse(g.Param("id"))
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
	}

	task, err := c.repo.Done(g, id)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	g.JSON(http.StatusOK, task)
}
