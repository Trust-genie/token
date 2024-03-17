package controller

import (
	"github.com/gin-gonic/gin"
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

}

func (controller Tagscontroller) Save(ctx gin.Context) {

}

func (controller Tagscontroller) Update(ctx gin.Context) {

}

func (controller Tagscontroller) Find(ctx gin.Context) {

}

func (controller Tagscontroller) Find_all(ctx gin.Context) {

}

func (controller Tagscontroller) Delete(ctx gin.Context) {

}
