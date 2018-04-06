package main
import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../proto"
	"google.golang.org/grpc/reflection"
	//"fmt"
	"fmt"
	"../../../../dao/factory"
	"../../../../utilities"
	"encoding/json"
)

const (
	port = ":66666"
)

type server struct{}

func (s *server) FindAllRadicacion(ctx context.Context, in *pb.GNricRequest) (*pb.GNricResponse, error) {
	if len(in.Name) > 0  {
		fmt.Print("se recibio un parametro "+in.Name)
		return &pb.GNricResponse{Message: "" + in.Name}, nil
	}
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return nil,nil
	}
	radicacionDao := factory.FactoryDao(config.Engine)
	radicaResult ,_ := radicacionDao.GetAll()
	b, err := json.Marshal(radicaResult)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &pb.GNricResponse{Message: string(b)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}