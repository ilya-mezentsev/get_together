package models

import "time"

type (
  MeetingMessage struct {
    Text string `db:"text"`
    SendingTime time.Time `db:"sending_time"`
    SenderId uint `db:"sender_id"`
    MeetingId uint `db:"meeting_id"`
  }

  MeetingRequestMessage struct {
    Text string `db:"text"`
    SendingTime time.Time `db:"sending_time"`
    SenderId uint `db:"sender_id"`
    AdminId uint `db:"admin_id"`
  }
)
