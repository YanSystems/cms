package repositories

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *ContentRepository) DeleteContent(coll string, id string) (string, error) {
	slog.Debug("DeleteContent called", "collection", coll, "id", id)

	_, err := r.DB.Collection(coll).DeleteOne(
		context.TODO(),
		bson.D{{Key: "id", Value: id}},
	)

	if err != nil {
		slog.Error("Failed to delete content", "collection", coll, "id", id, "error", err)
		return "", err
	}

	slog.Info("Content deleted successfully", "collection", coll, "id", id)
	return id, nil
}

func (r *ContentRepository) DeleteClass(coll string, class string) ([]string, error) {
	slog.Debug("DeleteClass called", "collection", coll, "class", class)

	contents, err := r.GetClass(coll, class)
	if err != nil {
		slog.Error("Failed to get class contents", "collection", coll, "class", class, "error", err)
		return nil, err
	}

	var ids []string
	for _, content := range contents {
		ids = append(ids, content.Id)
	}
	slog.Debug("Class contents to be deleted", "collection", coll, "class", class, "ids", ids)

	_, err = r.DB.Collection(coll).DeleteMany(
		context.TODO(),
		bson.D{{Key: "class", Value: class}},
	)

	if err != nil {
		slog.Error("Failed to delete class contents", "collection", coll, "class", class, "error", err)
		return nil, err
	}

	slog.Info("Class contents deleted successfully", "collection", coll, "class", class)
	return ids, nil
}

func (r *ContentRepository) DeleteCollection(coll string) ([]string, error) {
	slog.Debug("DeleteCollection called", "collection", coll)

	contents, err := r.GetCollection(coll)
	if err != nil {
		slog.Error("Failed to get collection contents", "collection", coll, "error", err)
		return nil, err
	}

	var ids []string
	for _, content := range contents {
		r.DeleteContent(coll, content.Id)
		ids = append(ids, content.Id)
	}
	slog.Debug("Collection contents to be deleted", "collection", coll, "ids", ids)

	err = r.DB.Collection(coll).Drop(context.TODO())
	if err != nil {
		slog.Error("Failed to drop collection", "collection", coll, "error", err)
		return nil, err
	}

	slog.Info("Collection dropped successfully", "collection", coll)
	return ids, nil
}
