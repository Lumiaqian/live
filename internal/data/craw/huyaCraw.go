package craw

import (
	"context"
	"live/internal/model"
)

type HuYaCraw interface {
	GetLiveUrl(ctx context.Context, roomId string) (*model.HuYaRoom, error)
}
