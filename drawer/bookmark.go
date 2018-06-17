package drawer

type BookmarkObject struct {
	id      *int
	url     *string
	title   *string
	comment *string
	tags    *[]string
}

func NewBookmark(id *int, url *string, title *string, comment *string, tags *[]string) *BookmarkObject {
	return &BookmarkObject{
		id:      id,
		title:   title,
		url:     url,
		comment: comment,
		tags:    tags,
	}
}

func (b *BookmarkObject) getTitle() *string {
	return b.title
}

func (b *BookmarkObject) getURL() *string {
	return b.url
}

func (b *BookmarkObject) getComment() *string {
	return b.comment
}

func (b *BookmarkObject) getTags() *[]string {
	return b.tags
}
