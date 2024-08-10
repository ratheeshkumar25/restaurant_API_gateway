package admin

import (
	"log"

	adminpb "github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/admin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (adminpb.AdminServiceClient, error) {
	grpc, err := grpc.NewClient(":8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: 8084 ")
		return nil, err
	}
	log.Printf("succefully connected to admin server at port :8084")
	return adminpb.NewAdminServiceClient(grpc), nil
}
