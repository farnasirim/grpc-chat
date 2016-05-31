package server

import (
	"fmt"
	"io"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/colonelmo/grpc-chat/proto"
	"github.com/colonelmo/grpc-chat/utils"
)

var usersLock = &sync.Mutex{}

var usersMap = make(map[string]chan pb.Message, 100)

type chatServer struct{}

func newChatServer() *chatServer {
	return &chatServer{}
}

func addListener(name string, msgQ chan pb.Message) {
	usersLock.Lock()
	defer usersLock.Unlock()
	usersMap[name] = msgQ
}

func removeListener(name string) {
	usersLock.Lock()
	defer usersLock.Unlock()
	delete(usersMap, name)
}

func hasListener(name string) bool {
	usersLock.Lock()
	defer usersLock.Unlock()
	_, exists := usersMap[name]
	return exists
}

func broadcast(sender string, msg pb.Message) {
	usersLock.Lock()
	defer usersLock.Unlock()
	for user, q := range usersMap {
		if user != sender {
			q <- msg
		}
	}
}

func listenToClient(stream pb.Chat_TransferMessageServer, messages chan<- pb.Message) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// ?
		}
		if err != nil {
			// ??
		}
		messages <- *msg
	}
}

func (s *chatServer) TransferMessage(stream pb.Chat_TransferMessageServer) error {
	clientIdentification, err := stream.Recv()
	var clientName string
	clientMailbox := make(chan pb.Message, 100)

	if err != nil {
		return err
	}
	if clientIdentification.Register {
		clientName = clientIdentification.Sender
		if hasListener(clientName) {
			return fmt.Errorf("name already exists")
		}
		addListener(clientName, clientMailbox)
	} else {
		return fmt.Errorf("need to register first")
	}

	clientMessages := make(chan pb.Message, 100)
	go listenToClient(stream, clientMessages)

	for {
		select {
		case messageFromClient := <-clientMessages:
			broadcast(clientName, messageFromClient)
		case messageFromOthers := <-clientMailbox:
			stream.Send(&messageFromOthers)
		}
	}
}

func Serve(address string, secure bool) error {

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	if secure {
		creds, err := credentials.NewServerTLSFromFile(utils.ConfigString("TLS_CERT"), utils.ConfigString("TLS_KEY"))
		if err != nil {
			return err
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterChatServer(grpcServer, newChatServer())
	grpcServer.Serve(lis)
	return nil
}
