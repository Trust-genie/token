package service

import (
	"github.com/go-playground/validator/v10"
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

