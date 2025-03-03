// flags/flag_es.go
package flags

import (
	"Blog-Server/models"
	"Blog-Server/service/es_service"
)

func EsIndex() {
	article := models.ArticleModel{}
	es_service.CreateIndexV2(article.Index(), article.Mapping())
}
