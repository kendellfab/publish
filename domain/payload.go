package domain

type PayloadRepo interface {
	GetPayload() *Payload
}

type Payload struct {
	Config      Config
	Categories  []*Category
	RecentPosts []*Post
}
