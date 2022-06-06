package section

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
)

type SectionController struct {
	service sections.Service
}

func NewSection(s sections.Service) *SectionController {
	return &SectionController{
		service: s,
	}
}

func (c SectionController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": s})
	}
}

func (c SectionController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		s, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, s)
	}
}

func (c *SectionController) Add() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		body := ctx.Request.Body
		data, _ := ioutil.ReadAll(body)

		var obj sections.Section

		sect := reflect.TypeOf(obj)
		for i := 0; i < sect.NumField(); i++ {
			field := sect.Field(i)
			if !strings.Contains(string(data), field.Tag.Get("json")) && field.Tag.Get("json") != "id" {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "missing field"})
				return
			}
		}

		if err := ctx.ShouldBindJSON(&obj); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := c.service.LastID()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id++

		obj.ID = id

		has, err := c.service.HasSectionNumber(obj.SectionNumber)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if has {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		s, err := c.service.Add(obj)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, s)
	}
}

func (c *SectionController) UpdateById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var obj sections.Section

		if err := ctx.ShouldBindJSON(&obj); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		s, err := c.service.UpdateById(id, obj)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, s)
	}
}

func (c *SectionController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}
