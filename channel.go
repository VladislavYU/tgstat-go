package tgstat_go

type ChannelResponse struct {
	Id                int         `json:"id"`
	Link              string      `json:"link"`
	Username          string      `json:"username"`
	Title             string      `json:"title"`
	About             string      `json:"about"`
	Category          string      `json:"category"`
	Country           string      `json:"country"`
	Language          string      `json:"Language"`
	Image100          string      `json:"image100"`
	Image640          string      `json:"image640"`
	ParticipantsCount int         `json:"participants_count"`
	TGStatRestriction interface{} `json:"tgstat_restrictions"`
}

type TGStatRestrictions struct {
	RedLabel   bool `json:"red_label"`
	BlackLabel bool `json:"black_label"`
}

type ChannelResponseResult struct {
	Status   string          `json:"status"`
	Response ChannelResponse `json:"response"`
}

type ChannelSearchResult struct {
	Status   string        `json:"status"`
	Response ChannelSearch `json:"response"`
}

type ChannelSearchItem struct {
	Id                int    `json:"id"`
	Link              string `json:"link"`
	Username          string `json:"username"`
	Title             string `json:"title"`
	About             string `json:"about"`
	Image100          string `json:"image100"`
	Image640          string `json:"image640"`
	ParticipantsCount int    `json:"participants_count"`
}

type ChannelSearch struct {
	Count int                 `json:"count"`
	Items []ChannelSearchItem `json:"items"`
}

type ChannelStatResult struct {
	Status   string              `json:"status"`
	Response ChannelStatResponse `json:"response"`
}

type ChannelStatResponse struct {
	ID                      int     `json:"id"`
	Title                   string  `json:"title"`
	Username                string  `json:"username"`
	PeerType                string  `json:"peer_type"`
	ParticipantCount        int     `json:"participants_count"`
	AvgPostReach            int     `json:"avg_post_reach"`
	AdvPostReach12h         int     `json:"adv_post_reach_12h"`
	AdvPostReach24h         int     `json:"adv_post_reach_24h"`
	AdvPostReach48h         int     `json:"adv_post_reach_48h"`
	ErrPercent              float64 `json:"err_percent"`
	Err24Percent            float64 `json:"err24_percent"`
	ErPercent               float64 `json:"er_percent"`
	DailyReach              int     `json:"daily_reach"`
	CiIndex                 float64 `json:"ci_index"`
	MentiosCount            int     `json:"mentions_count"`
	ForwardCount            int     `json:"forwards_count"`
	MentioningChannelsCount int     `json:"mentioning_channels_count"`
}

type ChannelPostsWithChannelResponseItem struct {
	ID            int64        `json:"id"`
	Date          int          `json:"date"`
	Views         int          `json:"views"`
	Link          string       `json:"link"`
	ChannelID     int          `json:"channel_id"`
	ForwardedFrom interface{}  `json:"forwarded_from"`
	IsDeleted     int          `json:"is_deleted"`
	Text          string       `json:"text"`
	Media         ChannelMedia `json:"media"`
}

type ChannelMedia struct {
	MediaType string `json:"media_type"`
	MimeType  string `json:"mime_type"`
	Size      int    `json:"size"`
}

type Channel struct {
	ID                int    `json:"id"`
	Link              string `json:"link"`
	Username          string `json:"username"`
	Title             string `json:"title"`
	About             string `json:"about"`
	Image100          string `json:"image100"`
	Image640          string `json:"image640"`
	ParticipantsCount int    `json:"participants_count"`
}

type ChannelPostsWithChannelResponse struct {
	Count      int                                   `json:"count"`
	TotalCount int                                   `json:"total_count"`
	Channel    Channel                               `json:"channel"`
	Items      []ChannelPostsWithChannelResponseItem `json:"items"`
}

type ChannelPostsWithChannelResult struct {
	Status   string                          `json:"status"`
	Response ChannelPostsWithChannelResponse `json:"response"`
}

type ChannelPostsResponseItem struct {
	ID            int64        `json:"id"`
	Date          int          `json:"date"`
	Views         int          `json:"views"`
	Link          string       `json:"link"`
	ChannelID     int          `json:"channel_id"`
	ForwardedFrom string       `json:"forwarded_from"`
	IsDeleted     int          `json:"is_deleted"`
	Text          string       `json:"text"`
	Media         ChannelMedia `json:"media"`
}

type ChannelPostsResponse struct {
	Count      int                        `json:"count"`
	TotalCount int                        `json:"total_count"`
	Channel    Channel                    `json:"channel"`
	Items      []ChannelPostsResponseItem `json:"items"`
}

type ChannelPostsResult struct {
	Status   string               `json:"status"`
	Response ChannelPostsResponse `json:"response"`
}

type MentionItem struct {
	MentionID   int    `json:"mentionId"`
	MentionType string `json:"mentionType"`
	PostID      int64  `json:"postId"`
	PostLink    string `json:"postLink"`
	PostDate    int    `json:"postDate"`
	ChannelID   int    `json:"channelId"`
}

type ChannelMentionsResponse struct {
	Items []MentionItem `json:"items"`
}

type ChannelMentionsResponseExtended struct {
	Items    []MentionItem `json:"items"`
	Channels []Channel     `json:"channels"`
}

type ChannelMentionsResult struct {
	Status   string                  `json:"status"`
	Response ChannelMentionsResponse `json:"response"`
}

type ChannelMentionsExtended struct {
	Status   string                          `json:"status"`
	Response ChannelMentionsResponseExtended `json:"response"`
}

type ForwardItem struct {
	ForwardID int    `json:"forwardId"`
	PostID    int64  `json:"postId"`
	PostLink  string `json:"postLink"`
	PostDate  int    `json:"postDate"`
	ChannelID int    `json:"channelId"`
}

type ChannelForwardsResponseExtended struct {
	Items    []ForwardItem `json:"items"`
	Channels []Channel     `json:"channels"`
}

type ChannelForwardsExtended struct {
	Status   string                          `json:"status"`
	Response ChannelForwardsResponseExtended `json:"response"`
}

type ChannelForwardsResponse struct {
	Items []ForwardItem `json:"items"`
}

type ChannelForwards struct {
	Status   string                  `json:"status"`
	Response ChannelForwardsResponse `json:"response"`
}

type ChannelSubscribersResponse struct {
	Period            string `json:"period"`
	ParticipantsCount uint   `json:"participants_count,string"`
}

type ChannelSubscribers struct {
	Status   string                       `json:"status"`
	Response []ChannelSubscribersResponse `json:"response"`
}

type ViewItem struct {
	Period     string `json:"period"`
	ViewsCount int    `json:"views_count"`
}

type ChannelViewsResponse struct {
	Period     string  `json:"period"`
	ViewsCount float64 `json:"views_count"`
}

type ChannelViews struct {
	Status   string                 `json:"status"`
	Response []ChannelViewsResponse `json:"response"`
}

type ChannelAvgReachResponse struct {
	Period        string  `json:"period"`
	AvgPostsReach float64 `json:"avg_posts_reach"`
}

type ChannelAvgReach struct {
	Status   string                    `json:"status"`
	Response []ChannelAvgReachResponse `json:"response"`
}

type ChannelErrResponse struct {
	Period string  `json:"period"`
	Err    float64 `json:"err"`
}

type ChannelErr struct {
	Status   string               `json:"status"`
	Response []ChannelErrResponse `json:"response"`
}

type ChannelAddPending struct {
	Status string `json:"status"`
}

type ChannelAddSuccess struct {
	Status   string `json:"status"`
	Response struct {
		ChannelId int `json:"channelId"`
	} `json:"response,omitempty"`
}
