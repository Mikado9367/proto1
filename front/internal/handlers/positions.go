package handlers

import (
	"errors"
	"net/http"

	"front/internal/clientgrpc/handlers"
	"front/internal/clientgrpc/secpos"
	"front/internal/models/api"
	"front/internal/tool"

	"github.com/gin-gonic/gin"
)

// c'est un peu caca, on devrait peut être avoir une couche abstraite pour ne rien changer si on passe
// de gRPC à autre chose.
// type ClientGRPC struct {
// 	pb.SecurityPosisionSettlementServiceClient
// }

// var myClient ClientGRPC

// type SecPosRequest struct {
// 	SearchCriteria ApiSearchCriteria
// }

// regarder cela : https://gin-gonic.com/docs/examples/bind-query-or-post/
// et pour le binding et controle : https://gin-gonic.com/docs/examples/bind-uri/

// http://localhost:8080/api/settlement/v1/security/welcome
func Welcome(c *gin.Context) {

	c.JSON(http.StatusOK, "welcome security")

	// appeler le service rpc dans un channel ?

}

// exemple : http://localhost:8080/positions/isin/12/account/-/restrictiontype/-?date=2024-10-22&phase=SOD

//http://localhost:8080/api/settlement/v1/security/position?businessperiodtype=SOD&businessdate=2024-10-07

// il y a deux choses :
// le binding de l'uri, mais aussi le bing des query string (filtres) .
// pour le moment, le binding des filtres ne fonctionne pas.
// pour le binding uri : https://gin-gonic.com/docs/examples/bind-uri/

func CheckAndValidate(c *gin.Context) (apiData api.ApiData, err error) {

	if err = c.BindQuery(&apiData.SearchFilter); err != nil {
		//		c.JSON(400, gin.H{"msg": err.Error()})
		return apiData, err
	}

	if err = c.ShouldBindUri(&apiData.SecurityPositionKey); err != nil {

		//		c.JSON(400, gin.H{"msg": err.Error()})
		return apiData, err
	}
	if len(apiData.SearchFilter.BusinessDate) > 0 {

		if !tool.IsDateValue(apiData.SearchFilter.BusinessDate) {
			// c.JSON(400, gin.H{"msg": "filter date is not valid"})
			err = errors.New("filter date is not valid")
			return apiData, err
		}
	}

	//FIXME: http://localhost:8080/api/settlement/v1/security/position/clientbic/12345/isin/22/account/111/restrictiontype/33?businessdate=2024-10-01&phase=ZOB
	//TODO: ajouter le controle du filtre pour verifier le domaine de valeur vs enum du fichier proto

	if len(apiData.SearchFilter.BusinessPeriodType) > 0 {

		_, ok := secpos.BusinessPeriodType_value[apiData.SearchFilter.BusinessPeriodType]
		if !ok {
			err = errors.New("period filter invalid")
			return apiData, err
		}
	}

	return apiData, nil
}

func GetOneSecurityPosition(c *gin.Context) {

	MyData, err := CheckAndValidate(c)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	feedback, err := handlers.GetOneSecurityPosition(MyData)

	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// abnormal situation if we get back more than onePosition with this function
	if len(feedback) > 1 {
		err = errors.Join(err, errors.New("unespect data with multiple values"))
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	// abnormal situation if we get back more than onePosition with this function
	if len(feedback) == 0 {
		err = errors.Join(err, errors.New("feedback getone empty"))
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, feedback[0])

	//	c.JSON(200, gin.H{"clientbic": MyData.SecurityPositionKey.Clientbic, "isin": MyData.SecurityPositionKey.Isin, "account": MyData.SecurityPositionKey.Account, "restrictiontype": MyData.SecurityPositionKey.Restrictiontype, "filter_date": MyData.SearchFilter.BusinessDate, "period": MyData.SearchFilter.BusinessPeriodType})
}

// cli := pb.NewSecurityPosisionSettlementServiceClient(config.ConnGRPC)

// answer := handlers.GetOneSecurityPosition(cli)

// log.Println(answer.SecurityPosition[0].SecurityPositionValue.PositionQuantity)

// recuper les infos puis ensuite, les convertir dans un format JSON pour ensuite, si pas d'erreur, répoindre un code 200 + contenu des datas.
// answer.SecurityPosition

// implémenter l'appel gRPC  avec un channel.
