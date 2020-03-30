package repositories

const (
	MeetingIdWithoutMeetingChat = 3
	MeetingType                 = "meeting"
	ArchivedStatus              = "archived"
)

var (
	NotExistsChatId = uint(len(MeetingChats) + 1)
)
