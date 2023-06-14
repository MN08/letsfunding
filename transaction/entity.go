package transaction

import (
	"letsfunding/campaign"
	"letsfunding/user"
	"time"
)

type (
	Transaction struct {
		ID         int
		CampaignID int
		UserID     int
		Amount     int
		Status     string
		Code       string
		CreatedAt  time.Time
		UpdatedAt  time.Time
		User       user.User
		Campaign   campaign.Campaign
	}
)
