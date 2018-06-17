package drawer

type BookmarkObject struct {
	id      int
	url     string
	title   string
	comment string
	tags    []string
}

func NewBookmark(id int, url string, title string, comment string, tags []string) *BookmarkObject {
	return &BookmarkObject{
		id:      id,
		title:   title,
		url:     url,
		comment: comment,
		tags:    tags,
	}
}

func (b *BookmarkObject) GetTitle() string {
	return b.title
}

func (b *BookmarkObject) GetURL() string {
	return b.url
}

func (b *BookmarkObject) GetComment() string {
	return b.comment
}

func (b *BookmarkObject) GetTags() []string {
	return b.tags
}
