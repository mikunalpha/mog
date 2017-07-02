package posts

import "github.com/mikunalpha/mog/api/shared/store"

type PostsData struct {
	Data []*store.Post `json:"data"`
}

type PostData struct {
	Data *store.Post `json:"data"`
}
