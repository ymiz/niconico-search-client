package common

type Response struct {
	Meta Meta      `json:"meta"`
	Data []Content `json:"data"`
}

type Meta struct {
	Status       int    `json:"status"`
	TotalCount   int    `json:"totalCount"`
	Id           string `json:"id"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type Content struct {
	ContentId              string `json:"contentId,omitempty"`
	Title                  string `json:"title,omitempty"`
	Description            string `json:"description,omitempty"`
	UserId                 int    `json:"userId,omitempty"`
	ViewCounter            int    `json:"viewCounter,omitempty"`
	MylistCounter          int    `json:"mylistCounter,omitempty"`
	LengthSeconds          int    `json:"lengthSeconds,omitempty"`
	ThumbnailUrl           string `json:"thumbnailUrl,omitempty"`
	StartTime              string `json:"startTime,omitempty"`
	ThreadId               int    `json:"threadId,omitempty"`
	CommentCounter         int    `json:"commentCounter,omitempty"`
	LastCommentTime        string `json:"lastCommentTime,omitempty"`
	CategoryTags           string `json:"categoryTags,omitempty"`
	ChannelId              int    `json:"channelId,omitempty"`
	Tags                   string `json:"tags,omitempty"`
	TagsExact              string `json:"tagsExact,omitempty"`
	LockTagsExact          string `json:"lockTagsExact,omitempty"`
	Genre                  string `json:"genre,omitempty"`
	GenreKeyword           string `json:"genre.keyword,omitempty"`
	CommunityId            int    `json:"communityId,omitempty"`
	ProviderType           string `json:"providerType,omitempty"`
	OpenTime               string `json:"openTime,omitempty"`
	LiveEndTime            string `json:"liveEndTime,omitempty"`
	TimeshiftEnabled       bool   `json:"timeshiftEnabled,omitempty"`
	ScoreTimeshiftReserved int    `json:"scoreTimeshiftReserved,omitempty"`
	CommunityText          string `json:"communityText,omitempty"`
	CommunityIcon          string `json:"communityIcon,omitempty"`
	MemberOnly             bool   `json:"memberOnly,omitempty"`
	LiveStatus             string `json:"liveStatus,omitempty"`
}
