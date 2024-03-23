package repository

import (
	"errors"

	request "github.com/liezner/token/data/requests"
	"github.com/liezner/token/helper"
	"github.com/liezner/token/models"
	"gorm.io/gorm"
)

type Tagsrepositoryimpl struct {
	DB *gorm.DB
}

func NewTagsRepo(Db *gorm.DB) Tagsrepository {
	return &Tagsrepositoryimpl{DB: Db}
}

// Delete implements Tagsrepository
func (t *Tagsrepositoryimpl) Delete(tagsId int) {
	var tags models.Tags
	result := t.DB.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// Implements the findall handler
func (t *Tagsrepositoryimpl) Find_all() []models.Tags {
	var tags models.Tags
	result := t.DB.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
	// this should return a slice not a var

}

// implements the Find by id ie Find func

func (t *Tagsrepositoryimpl) Find(tagsId int) (tags models.Tags, err error) {
	// in the tutorial they explictly defined tags again but the compiler says i can discard that
	result := t.DB.Find(&tags, tagsId)

	if result != nil {
		return tags, nil
	} else {
		return tags, errors.New("ID not Found")
	}
}

// The save function
// This function is implementing the `Save` method for the `Tagsrepositoryimpl` struct. It takes a
// `models.Tags` object as a parameter, attempts to create a new record in the database using the
// `Create` method provided by the `gorm.DB` instance stored in the `DB` field of the
// `Tagsrepositoryimpl` struct, and then checks for any errors that occurred during the operation using
// the `helper.ErrorPanic` function.
func (t *Tagsrepositoryimpl) Save(tags models.Tags) {

	result := t.DB.Create(&tags)
	helper.ErrorPanic(result.Error)
}

// the update fuction
func (t *Tagsrepositoryimpl) Update(tag models.Tags) {

	var updatetag = request.Updatedatarequest{
		Uuid: tag.Salary,
		Name: tag.Name_first,
	}
	result := t.DB.Model(&tag).Updates(updatetag)
	helper.ErrorPanic(result.Error)
}
