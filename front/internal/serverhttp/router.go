package serverhttp

// ------------------------------------------------------------------------------
// router.Run(config.Envs.PublicHost + config.Envs.Port)
// Si on met localhost, ce con fait une erreur sur le protocol TCP en disant :
// listen tcp: address http://www.toto.com:8080: too many colons in address

// regarder cela : https://gin-gonic.com/docs/examples/bind-query-or-post/
// et pour le binding et controle : https://gin-gonic.com/docs/examples/bind-uri/
// ------------------------------------------------------------------------------

import (
	"front/config"
	"front/internal/handlers"

	"github.com/gin-gonic/gin"
)

// setup log
// on peut lire l'article : https://github.com/gin-gonic/gin/issues/1376

func InitMyRouter(router *gin.Engine) {

	sec := router.Group(config.Envs.APIrelease)

	security := sec.Group("/security")
	security.GET("/welcome", handlers.Welcome)

	// reprendre là, pourquoi url passe pas et fait 404
	// http://localhost:8080/api/settlement/v1/security/position/clientbic/12345/isin/22/account/11/restrictiontype/33

	// List of all endpoints and its function
	var urlGetOnePosition = "/position/clientbic/:clientbic/isin/:isin/account/:account/restrictiontype/:restrictiontype"
	security.GET(urlGetOnePosition, handlers.GetOneSecurityPosition)

	// start framework
	router.Run(config.Envs.Port)

}
