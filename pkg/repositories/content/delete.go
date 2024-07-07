package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *ContentRepository) DeleteContent(coll string, id string) (string, error) {
	_, err := r.DB.Collection(coll).DeleteOne(
		context.TODO(),
		bson.D{{Key: "id", Value: id}},
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *ContentRepository) DeleteClass(coll string, class string) ([]string, error) {
	contents, err := r.GetClass(coll, class)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, content := range contents {
		ids = append(ids, content.Id)
	}

	_, err = r.DB.Collection(coll).DeleteMany(
		context.TODO(),
		bson.D{{Key: "class", Value: class}},
	)

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *ContentRepository) DeleteCollection(coll string) ([]string, error) {

	contents, err := r.GetCollection(coll)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, content := range contents {
		r.DeleteContent(coll, content.Id)
		ids = append(ids, content.Id)
	}

	err = r.DB.Collection(coll).Drop(context.TODO())
	if err != nil {
		return nil, err
	}

	return ids, nil
}
