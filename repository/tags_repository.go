package repository

import (
	"github.com/liezner/token/models"
)

type Tagsrepository interface {
	Save(tags models.Tags)
	Update(tags models.Tags)
	Delete(tagsId int)
	Find(models.Tags) (tags models.Tags, err error)
	Find_all() []models.Tags
}
