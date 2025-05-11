package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type LinkModel struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM123456789")

func NewLink(url string) *LinkModel {
	return &LinkModel{
		Url:  url,
		Hash: RandString(6),
	}
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
