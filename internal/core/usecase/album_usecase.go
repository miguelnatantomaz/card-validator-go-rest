package usecase

import (
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/domain"
)

type AlbumUsecase struct {
	albums []domain.Album
}

func NewAlbumUsecase() *AlbumUsecase {
	return &AlbumUsecase{
		albums: []domain.Album{
			{ID: "1", Title: "Title 1", Artist: "Artist 1", Price: 234},
			{ID: "2", Title: "Title 2", Artist: "Artist 2", Price: 65},
		},
	}
}

func (u *AlbumUsecase) GetAllAlbums() []domain.Album {
	return u.albums
}

func (u *AlbumUsecase) GetAlbumByID(id string) *domain.Album {
	for _, a := range u.albums {
		if a.ID == id {
			return &a
		}
	}
	return nil
}

func (u *AlbumUsecase) AddAlbum(album domain.Album) {
	u.albums = append(u.albums, album)
}
