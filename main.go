package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	service "github.com/ymcagodme/shortn/service"
	"google.golang.org/grpc"

	pb "github.com/ymcagodme/shortn/proto"
)

const (
	rpcAddr = "localhost:52002"
)

func loggingDecorator(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL)
		f(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	rawurl := ""
	if url, ok := r.URL.Query()["url"]; ok {
		rawurl = url[0]
	}
	sid, err := service.AddPage(rawurl)
	if err != nil {
		log.Printf("[handler] error = %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s -> %s\n", sid, rawurl)
}

var rpcConn *grpc.ClientConn
var rpcClient pb.ShortnClient

func handlerUseRPC(w http.ResponseWriter, httpReq *http.Request) {
	if rpcConn == nil {
		dialCtx, dialCancel := context.WithTimeout(context.Background(), time.Second)
		defer dialCancel()
		conn, err := grpc.DialContext(dialCtx, rpcAddr, grpc.WithInsecure(), grpc.WithBlock())

		if err != nil {
			errMsg := fmt.Sprintf("failed to connect rpc service: %v", err)
			log.Printf(errMsg)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}
		rpcConn = conn
	}
	if rpcClient == nil {
		rpcClient = pb.NewShortnClient(rpcConn)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := rpcClient.AddPageRpc(ctx, &pb.AddPageRequest{RawUrl: "raw_url"})

	if err != nil {
		errMsg := fmt.Sprintf("failed to AddPageRequest: %v", err)
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "res: %v", r.GetShortUrl())
}

func main() {
	http.HandleFunc("/shortn", loggingDecorator(handler))
	http.HandleFunc("/shortn_rpc", loggingDecorator(handlerUseRPC))
	log.Printf("Server starts listening :8080")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
