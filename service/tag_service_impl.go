package service

import (
	"github.com/go-playground/validator/v10"
	request "github.com/liezner/token/data/requests"
	"github.com/liezner/token/data/response"
	"github.com/liezner/token/repository"
)

type Tagserviceimpl struct {
	Tagsrepository repository.Tagsrepository
	Validate       *validator.Validate
}

func NewTagServiceimpl(tagrepo repository.Tagsrepository, validate *validator.Validate) Tagservice {
	return &Tagserviceimpl{
		Tagsrepository: tagrepo,
		Validate:       validate,
	}
}

// the tagservice init
func (t *Tagserviceimpl) Create(tags request.Createdatarequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.name,
	}
	t.Tagsrepository.Save(tagModel)
}

// the delete tagservice

func (t, *Tagserviceimpl)Delete(tagsId int){
	t.Tagsrepository.Delete(tagsId)
}

func (t *Tagserviceimpl)Find_all() []response.Tagresponse{
	var tags []response.Tagresponse

	result := t.Tagsrepository.Find_all()

	for _,value := range result{
		tag := response.Tagresponse(
			Id: value.id,
			Name: value.name,
		)
		tags = append(tags, tag)

	}
	return tags
}

// Find by Id implementation
func (t *Tagserviceimpl)Find(tagId int)response.Tagresponse{
	tagData, err := t.Tagsrepository.Find(tagId)
	helper.ErrorPanic(err)

	tagresponse := response.Tagresponse{
		Id: tagData.ID,
		Name: tagData.Name
	}
	return tagresponse
}

// the update implementation
func(t *Tagserviceimpl)Update(tagId	int)response.Tagresponse{
	tagData, err := t.Tagsrepository.Find(tagId)
	helper.ErrorPanic(err)

	tagData.Name = tags.Name
	t.Tagsrepository.Update(tagData)
}