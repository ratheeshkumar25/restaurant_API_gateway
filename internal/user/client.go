package user

import (
	"log"

	userpb "github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (userpb.UserServicesClient, error) {
	grpc, err := grpc.NewClient(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing client on port:8082")
		return nil, err
	}
	log.Printf("Successfully connected to client on port 8082")
	return userpb.NewUserServicesClient(grpc), err
}
