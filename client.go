package main

import (
	"context"
	"flag"
	"log"

	"calculator-client/api"

	"google.golang.org/grpc"
)

func main() {
	num1Ptr := flag.Int("num1", 0, "Number 1")
	num2Ptr := flag.Int("num2", 1, "Number 2")
	oprPtr := flag.String("opr", "+", "Calculator Operator (+,-,*,/)")

	serverHostPtr := flag.String("srv", "localhost:9000", "Server host (IP_ADDR_SRV:PORT")

	flag.Parse()
	var conn *grpc.ClientConn

	conn, errConn := grpc.Dial(*serverHostPtr, grpc.WithInsecure())
	if errConn != nil {
		log.Fatal("did not connect : ", errConn)
	}

	defer conn.Close()

	c := api.NewCalculatorClient(conn)

	respon, err := c.DoCalc(context.Background(), &api.CalculatorInputMessage{
		Num1: int32(*num1Ptr), Num2: int32(*num2Ptr), Operator: *oprPtr})

	if err != nil {
		log.Fatal("Error when call DO Calc", err)
	}
	// log.Print(respon)
	// go run client.go -srv localhost:9000 -num1 2 -num2 3 -opr "+"
	log.Printf("%v", respon.ResNum)
}
