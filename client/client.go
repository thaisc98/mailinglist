package main

import (
	"context"
	"log"
	pb "mailinglist/proto"
	"time"
)

func logResponse(res *pb.EmailResponse, err error){
	if err != nil {
		log.Fatalf(" Error: %v", err)
	}
	if res.EmailEntry == nil {
		log.Printf(" email not found")
	} else {
		log.Printf(" response: %v", res.EmailEntry)	
	}
}

func createEmail(client pb.MailingListServiceClient, addr string) {
	log.Println("create email") 

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateEmail(ctx, &pb.CreateEmailRequest{EmailAddr: addr})
	logResponse(res, err)
}

func getEmail(client pb.MailingListServiceClient, addr string){
	log.Println("get email") 

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetEmail(ctx, &pb.GetEmailRequest{EmailAddr: addr})
	logResponse(res, err)
}