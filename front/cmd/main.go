package main

import (
	"fmt"
	"front/config"
	"front/internal/serverhttp"

	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// var addrServerGRPC string = "0.0.0.0:50051"
// var ConnGRPC *grpc.ClientConn

// mettre en place l'utilisation de la log gin
// regarder les bind validator de gin, pour peut être ne pas appeler gRPC si de la merde rempli
// exemple là : https://github.com/gin-gonic/examples/blob/master/custom-validation/server.go
// et là : https://github.com/gin-gonic/examples/blob/master/custom-validation/server.go

func main() {


// 21/10/24 : ajouter un error handler via le middleware de gin si j'ai bien compris
// on peut commencer avec https://stackoverflow.com/questions/69948784/how-to-handle-errors-in-gin-middleware



	// setup gRPC
	var err error
	config.ConnGRPC, err = grpc.NewClient(config.AddrServerGRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		gin.DefaultErrorWriter.Write([]byte(fmt.Sprintf("Failed to connect : %v\n", err)))
		log.Fatalf("Failed to connect : %v\n", err)
	}

	// setup http server for Rest Api purpose
	router := gin.Default()
	serverhttp.InitMyRouter(router)

	// log.Printf(("start Settlement gRPC client service... %s\n"), addrServerGRPC)
	// c := pb.NewSecurityPosisionSettlementServiceClient(conn)
	defer config.ConnGRPC.Close()
	defer log.Printf(("closing... %s\n"), config.AddrServerGRPC)

}
