package fake

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"github.com/slipneff/notes/internal/pkg/models"
)

func Note() models.Note {
	return models.Note{
		Id:      uuid.New(),
		Header:  randomdata.Noun(),
		Content: randomdata.Paragraph(),
	}
}
