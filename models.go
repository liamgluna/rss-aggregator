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

func databaseFeedToFeed(dbFeedFollow database.Feed) Feed {
	return Feed{
		dbFeedFollow.ID,
		dbFeedFollow.CreatedAt,
		dbFeedFollow.UpdatedAt,
		dbFeedFollow.Name,
		dbFeedFollow.Url,
		dbFeedFollow.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}

	return feeds

}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}


func databaseFeedFollowToFeedFollow(dbUse database.FeedFollow) FeedFollow {
	return FeedFollow{
		dbUse.ID,
		dbUse.CreatedAt,
		dbUse.UpdatedAt,
		dbUse.UserID,
		dbUse.FeedID,
	}

}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollow := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollow = append(feedFollow, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}

	return feedFollow

}