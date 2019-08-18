package targets

type Field int

func (f Field) String() string {
	switch f {
	case ContentId:
		return "contentId"
	case Title:
		return "title"
	case Description:
		return "description"
	case UserId:
		return "userId"
	case ChannelId:
		return "channelId"
	case CommunityId:
		return "communityId"
	case ProviderType:
		return "providerType"
	case Tags:
		return "tags"
	case TagsExact:
		return "tagsExact"
	case CategoryTags:
		return "categoryTags"
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
	case TimeshiftEnabled:
		return "timeshiftEnabled"
	case ScoreTimeshiftReserved:
		return "scoreTimeshiftReserved"
	case ThumbnailUrl:
		return "thumbnailUrl"
	case CommunityText:
		return "communityText"
	case CommunityIcon:
		return "communityIcon"
	case MemberOnly:
		return "memberOnly"
	case LiveStatus:
		return "liveStatus"
	default:
		return "unknown"
	}
}

const (
	Unknown Field = iota
	ContentId
	Title
	Description
	UserId
	ChannelId
	CommunityId
	ProviderType
	Tags
	TagsExact
	CategoryTags
	ViewCounter
	CommentCounter
	OpenTime
	StartTime
	LiveEndTime
	TimeshiftEnabled
	ScoreTimeshiftReserved
	ThumbnailUrl
	CommunityText
	CommunityIcon
	MemberOnly
	LiveStatus
)
