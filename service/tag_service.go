package service

import (
	request "github.com/liezner/token/data/requests"
	"github.com/liezner/token/data/response"
)

type Tagservice interface {
	Create(tag request.Createdatarequest)
	Update(tag request.Updatedatarequest)
	Delete(tagid int)
	Find(tagid int) response.Tagresponse
	Find_all() []response.Tagresponse
}
