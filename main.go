package main

import (
	"fmt"
	"sync"

	pb "github.com/SemyonHoyrish/grpc-example/routes/user"
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
	savedFeatures []*pb.Feature // read-only after initialized

	mu         sync.Mutex // protects routeNotes
	routeNotes map[string][]*pb.RouteNote
}

func main() {
	fmt.Println("TEST")
}
