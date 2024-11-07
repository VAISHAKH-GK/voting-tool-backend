package poll

import (
	"context"

	"github.com/VAISHAKH-GK/voting-tool-backend/pkg/mongodb"
)

type Option struct {
	Name string `json:"name" form:"name" bson:"name"`
	Vote int    `json:"vote" form:"vote" bson:"vote"`
}

type Poll struct {
	Id          string   `json:"id" form:"id" bson:"_id"`
	CreatedBy   string   `json:"createdBy" form:"createdBy" bson:"createdBy"`
	Name        string   `json:"name" form:"name" bson:"name"`
	Description string   `json:"description" form:"description" bson:"description"`
	Options     []Option `json:"options" form:"options" bson:"options"`
	Total       int      `json:"total" form:"total" bson:"total"`
}

func (p *Poll) AddPoll() {
	mongodb.Poll().InsertOne(context.Background(), p)
}
