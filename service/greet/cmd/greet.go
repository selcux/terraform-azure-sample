package main

import (
	"log"

	"github.com/selcux/terraform-azure-sample/service/greet"
)

func main() {
	svc := greet.NewService()
	err := svc.Run("", 9010)
	if err != nil {
		log.Fatalf("Unable to serve with GRPC %v", err)
	}
}
