package interfaces

import (
	"github.com/kendellfab/publish/domain"
)

type AuthorCache struct {
	*domain.LruCache
}

func NewAuthorCache(size int) (*AuthorCache, error) {
	cache, err := domain.NewLruCache(size)
	if err != nil {
		return nil, err
	}
	return &AuthorCache{LruCache: cache}, nil
}

func (ac *AuthorCache) Add(id string, user *domain.User) {
	ac.LruCache.Add(id, user)
}

func (ac *AuthorCache) Get(id string) (*domain.User, bool) {
	if val, ok := ac.LruCache.Get(id); ok {
		return val.(*domain.User), true
	}
	return nil, false
}

type CategoryCache struct {
	*domain.LruCache
}

func NewCategoryCache(size int) (*CategoryCache, error) {
	cache, err := domain.NewLruCache(size)
	if err != nil {
		return nil, err
	}
	return &CategoryCache{LruCache: cache}, nil
}

func (cc *CategoryCache) Add(id string, cat *domain.Category) {
	cc.LruCache.Add(id, cat)
}

func (cc *CategoryCache) Get(id string) (*domain.Category, bool) {
	if val, ok := cc.LruCache.Get(id); ok {
		return val.(*domain.Category), true
	}
	return nil, false
}

type PostCache struct {
	*domain.LruCache
}

func NewPostCache(size int) (*PostCache, error) {
	cache, err := domain.NewLruCache(size)
	if err != nil {
		return nil, err
	}
	return &PostCache{LruCache: cache}, nil
}

func (pc *PostCache) Add(id string, post *domain.Post) {
	pc.LruCache.Add(id, post)
}

func (pc *PostCache) Get(id string) (*domain.Post, bool) {
	if val, ok := pc.LruCache.Get(id); ok {
		return val.(*domain.Post), true
	}
	return nil, false
}

type PageCache struct {
	*domain.LruCache
}

func NewPageCache(size int) (*PageCache, error) {
	cache, err := domain.NewLruCache(size)
	if err != nil {
		return nil, err
	}
	return &PageCache{LruCache: cache}, nil
}

func (pc *PageCache) Add(id string, page *domain.Page) {
	pc.LruCache.Add(id, page)
}

func (pc *PageCache) Get(id string) (*domain.Page, bool) {
	if val, ok := pc.LruCache.Get(id); ok {
		return val.(*domain.Page), true
	}
	return nil, false
}
