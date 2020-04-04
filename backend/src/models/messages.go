package models

import "time"

type (
	Chat struct {
		Id        uint      `db:"id"`
		Type      string    `db:"type"`
		Status    string    `db:"status"`
		CreatedAt time.Time `db:"created_at"`
	}

	Message struct {
		ChatId      uint      `db:"chat_id"`
		Text        string    `db:"text"`
		SendingTime time.Time `db:"sending_time"`
		SenderId    uint      `db:"sender_id"`
		MeetingId   uint      `db:"meeting_id"`
	}
)
