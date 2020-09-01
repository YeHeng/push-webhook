package model

type TextMessage struct {
	MsgType             string   `json:"msgtype"`
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type MarkdownMessage struct {
	MsgType             string    `json:"msgtype"`
	Markdown            *Markdown `json:"markdown"`
	MentionedList       []string  `json:"mentioned_list"`
	MentionedMobileList []string  `json:"mentioned_mobile_list"`
}

type Markdown struct {
	Content string `json:"content"`
}

type NewsMessage struct {
	MsgType  string    `json:"msgtype"`
	Markdown *Markdown `json:"news"`
}

type News struct {
	Articles *[]Article `json:"articles"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}
