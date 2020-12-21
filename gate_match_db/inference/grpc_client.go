package inference

import (
	"context"
	"log"
	"os"
	"errors"

	InferenceMsg "github.com/AlcheraInc/InferenceMsg/InferenceService"

	"google.golang.org/grpc"
)

var inferenceClient InferenceMsg.InferenceServiceClient
var inferenceConn *grpc.ClientConn
var grpcContext context.Context

func ConnectInferenceService() error {
	/* 연결을 계속 가지고있으면 안될수도 있음 */
	var err error
	inference_server_address := os.Getenv("INFERENCE_SERVER_ADDRESS")
	inferenceConn, err = grpc.Dial(inference_server_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err)
		return err
	}

	inferenceClient = InferenceMsg.NewInferenceServiceClient(inferenceConn)
	grpcContext = context.Background()

	/* Status Check */
	status_res, err := inferenceClient.RunStatus(grpcContext, &InferenceMsg.Void{})
	log.Println(status_res.GetCode())
	log.Println(status_res.GetMsg())

	if status_res.GetCode() != "INF-00" {
		return errors.New("Inference server at" + inference_server_address + "not ready.")
	}

	return nil
}
