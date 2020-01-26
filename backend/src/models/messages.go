package models

import "time"

type (
  Chat struct {
    ID uint `db:"id"`
    Type string `db:"type"`
    Status string `db:"status"`
    CreatedAt time.Time `db:"created_at"`
  }

  Message struct {
    ChatId uint `db:"chat_id"`
    Text string `db:"text"`
    SendingTime time.Time `db:"sending_time"`
    SenderId uint `db:"sender_id"`
  }

  MeetingMessage struct {
    Message
    MeetingId uint `db:"meeting_id"`
  }

  MeetingRequestMessage struct {
    Message
    AdminId uint `db:"admin_id"`
  }
)
