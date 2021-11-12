package controller

import (
	"context"

	"github.com/sevenpok/api-rabbit/database"
	"github.com/sevenpok/api-rabbit/models"
)

var collection = database.GetCollection("games")
var ctx = context.Background()

func Create(game models.Game) error {
	var err error
	_, err = collection.InsertOne(ctx, game)

	if err != nil {
		return err
	}
	return nil
}
