package repository

import (
	"token/models"
)
type Tagsrepository interface {
	Save(tags models.Tags)
	Update(tags models.Tags)
	Delete(tagsId int)
	Find(tags models.Tags) (tags models.Tags, err error)
	Find_all() []tags models.Tags
}
