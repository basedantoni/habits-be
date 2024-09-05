package factory

import (
	"basedantoni/habits-be/entities"
	"fmt"
)

type EntityFactory interface {
    CreateEntity(entityType string) (entities.Entity, error)
}

type SimpleEntityFactory struct{}

func (f *SimpleEntityFactory) CreateEntity(entityType string) (entities.Entity, error) {
    switch entityType {
    case "habit":
        return &entities.Habit{}, nil
    case "contribution":
        return &entities.Contribution{}, nil
    default:
        return nil, fmt.Errorf("unknown entity type %s", entityType)
    }
}