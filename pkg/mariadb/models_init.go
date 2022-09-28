package mariadb

import (
	"whdsl/pkg/article"
)

func (b *Backend) initModels() error {
	models := []Model{
		new(article.Article),
	}

	for i := range models {
		err := models[i].init(ctx, b.bunDB)
		if err != nil {
			return err
		}
	}

}
