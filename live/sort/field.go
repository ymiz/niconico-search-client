package sort

type Field int

func (f Field) String() string {
	switch f {
	case UserId:
		return "userId"
	case ChannelId:
		return "channelId"
	case CommunityId:
		return "communityId"
	case ViewCounter:
		return "viewCounter"
	case CommentCounter:
		return "commentCounter"
	case OpenTime:
		return "openTime"
	case StartTime:
		return "startTime"
	case LiveEndTime:
		return "liveEndTime"
	case ScoreTimeshiftReserved:
		return "scoreTimeshiftReserved"
	default:
		return "Unknown"
	}
}

const (
	Unknown Field = iota
	UserId
	ChannelId
	CommunityId
	ViewCounter
	CommentCounter
	OpenTime
	StartTime
	LiveEndTime
	ScoreTimeshiftReserved
)
