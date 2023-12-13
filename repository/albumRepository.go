package repository

import (
	"fmt"
	"gogin/model"

	"gorm.io/gorm"
)

type AlbumRepository interface {
	FindAllAlbums() ([]model.Album, error)
	FindAlbumById(id int) (model.Album, error)
	Save(album *model.Album) (model.Album, error)
	DeleteAlbumById(id int) error
}

type AlbumRepositoryImpl struct {
	db *gorm.DB
}

func (a AlbumRepositoryImpl) FindAllAlbums() ([]model.Album, error) {
	var albums []model.Album

	var err = a.db.Find(&albums).Error
	if err != nil {
		fmt.Println("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return albums, nil
}

func (a AlbumRepositoryImpl) FindAlbumById(id int) (model.Album, error) {
	album := model.Album{
		Id: id,
	}
	err := a.db.First(&album).Error
	if err != nil {
		fmt.Println("Got and error when find album by id. Error: ", err)
		return model.Album{}, err
	}
	return album, nil
}

func (u AlbumRepositoryImpl) Save(album *model.Album) (model.Album, error) {
	var err = u.db.Save(album).Error
	if err != nil {
		fmt.Println("Got an error when save album. Error: ", err)
		return model.Album{}, err
	}
	return *album, nil
}

func (u AlbumRepositoryImpl) DeleteAlbumById(id int) error {
	err := u.db.Delete(&model.Album{}, id).Error
	if err != nil {
		fmt.Println("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

func AlbumRepositoryInit(db *gorm.DB) *AlbumRepositoryImpl {
	db.AutoMigrate(&model.Album{})
	return &AlbumRepositoryImpl{
		db: db,
	}
}
