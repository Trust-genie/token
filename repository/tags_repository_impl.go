package repository

import (
	"errors"

	request "github.com/liezner/token/data/requests"
	"github.com/liezner/token/models"
	"github.com/liezner/token/service"
	"gorm.io/gorm"
)

type Tagsrepositoryimpl struct {
	Db *gorm.DB
}

func NewTagsRepo(Db *gorm.Db) Tagsrepository {
	return &Tagsrepositoryimpl{Db: db}
}


// Delete implements Tagsrepository
func (t *Tagsrepositoryimpl) Delete(tagsId int) {
	var tags models.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// Implements the findall handler
func (t *Tagsrepositoryimpl) Find_all() []models.Tags {
	var tags []models.Tags

	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags

}

// implements the Find by id ie Find func

func (t *Tagsrepositoryimpl) Find(tagsId int) (tags models.Tags, err error) {
	// in the tutorial they explictly defined tags again but the compiler says i can discard that
	result := t.Db.Find(&tags, tagsId)

	if result != nil {
		return tags, nil
	} else {
		return tags, errors.New("ID not Found")
	}
}

// The save function
func (t *Tagsrepositoryimpl) Save(tags models.Tags) {
	var tags models.Tags

	result := t.Db.Create(&tags)
	helper.ErrorPanic(result)
}

// the update fuction
func (t *Tagsrepositoryimpl) Update(tag models.Tags) {

	var updatetag = request.Updatedatarequest{
		Uuid: tag.Id,
		Name: tag.Name,
	}
	result := t.Db.Model(&tag).Updates(updatetag)
	helper.ErrorPanic(result.Error)
}
