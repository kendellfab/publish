package domain

type PayloadRepo interface {
	GetPayload() *Payload
}

type Payload struct {
	Config      ConfigApp
	Categories  []*Category
	RecentPosts []*Post
	Pages       []*Page
}
