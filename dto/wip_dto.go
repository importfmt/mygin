package dto

import (
	"time"

	"mygin.com/mygin/model"
)

type WipDto struct {
	CreatedAt time.Time `json:"created_at"`
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:username`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Reply     string    `json:"reply"`
	Status    bool      `json:"status"`
}

func ToWipDto(wips model.Wips) WipDto {
	return WipDto{
		CreatedAt: wips.CreatedAt,
		UserID:    wips.UserID,
		Username:  wips.Username,
		ID:        wips.ID,
		Title:     wips.Title,
		Desc:      wips.Desc,
		Reply:     wips.Reply,
		Status:    wips.Status,
	}

}
