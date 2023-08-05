package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/liamgluna/rss-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUse database.User) User {
	return User{
		dbUse.ID,
		dbUse.CreatedAt,
		dbUse.UpdatedAt,
		dbUse.Name,
		dbUse.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbUse database.Feed) Feed {
	return Feed{
		dbUse.ID,
		dbUse.CreatedAt,
		dbUse.UpdatedAt,
		dbUse.Name,
		dbUse.Url,
		dbUse.UserID,
	}
}

func databaseFeedsToFeeds(dbUse []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbUse {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}

	return feeds

}
