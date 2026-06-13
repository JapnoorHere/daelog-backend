package services

import (
	"context"
	"time"

	"github.com/japnoor/daelog/internal/db"
	"github.com/japnoor/daelog/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type EventService struct {
	db *db.MongoDB
}

func NewEventService(db *db.MongoDB) *EventService {
	return &EventService{db: db}
}

func (s *EventService) Create(req *models.CreateEventRequest) (*models.Event, error) {
	event := &models.Event{
		ID:        bson.NewObjectID(),
		SessionID: req.SessionID,
		EventType: req.EventType,
		Payload:   req.Payload,
		Repo:      req.Repo,
		FilePath:  req.FilePath,
		Ts:        req.Ts,
		Synced:    false,
		CreatedAt: time.Now().UnixMilli(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.Database.Collection("events").InsertOne(ctx, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventService) GetByDate(from, to int64) ([]models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"ts": bson.M{
			"$gte": from,
			"$lte": to,
		},
	}

	cursor, err := s.db.Database.Collection("events").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []models.Event
	if err := cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}
