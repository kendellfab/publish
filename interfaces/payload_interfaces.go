package interfaces

import (
	"github.com/kendellfab/publish/domain"
	"sync"
)

type PayloadRepoImpl struct {
	config domain.Config
	cr     domain.CategoryRepo
	pr     domain.PostRepo
}

func NewPayloadRepo(config domain.Config, cr domain.CategoryRepo, pr domain.PostRepo) domain.PayloadRepo {
	repo := &PayloadRepoImpl{config: config, cr: cr, pr: pr}
	return repo
}

func (p *PayloadRepoImpl) GetPayload() *domain.Payload {
	pay := &domain.Payload{}
	pay.Config = p.config.AppConfig

	var wg sync.WaitGroup
	wg.Add(2)

	go func(wait *sync.WaitGroup) {
		defer wait.Done()
		if cats, catErr := p.cr.GetAllCount(); catErr == nil {
			pay.Categories = cats
		}
	}(&wg)

	go func(wait *sync.WaitGroup) {
		defer wait.Done()
		if posts, pErr := p.pr.FindPublished(0, 10); pErr == nil {
			pay.RecentPosts = posts
		}
	}(&wg)

	wg.Wait()
	return pay
}
