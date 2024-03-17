package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	request "github.com/liezner/token/data/requests"
	"github.com/liezner/token/data/response"
	"github.com/liezner/token/service"
)

type Tagscontroller struct {
	tagService service.Tagservice
}

func Newtagcontroller(service service.Tagservice) *Tagscontroller {
	return &Tagscontroller{
		tagService: service,
	}
}

func (controller Tagscontroller) Create(ctx gin.Context) {
	createtagsreq := request.Createdatarequest{}
	err := ctx.ShouldBindJSON(createtagsreq)
	helper.ErrorPanic(err)

	controller.tagService.Create(createtagsreq)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller Tagscontroller) Save(ctx gin.Context) {

}

func (controller Tagscontroller) Update(ctx gin.Context) {
	updatetagreq := request.Updatedatarequest{}
	err := ctx.ShouldBindJSON(&updatetagreq)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)

	updatetagreq.Id = id

	controller.tagService.Update(updatetagreq)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller Tagscontroller) Find(ctx gin.Context) {

}

func (controller Tagscontroller) Find_all(ctx gin.Context) {

}

func (controller Tagscontroller) Delete(ctx gin.Context) {

}
