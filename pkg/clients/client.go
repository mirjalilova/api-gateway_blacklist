package clients

import (
	"log/slog"

	"github.com/mirjalilova/api-gateway_blacklist/internal/config"
	pb "github.com/mirjalilova/api-gateway_blacklist/internal/genproto/black_list"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type Clients struct {
	AdminClient     pb.AdminServiceClient
	HrClient        pb.HRServiceClient
	BlacklistClient  pb.BlackListServiceClient

}

func NewClients(cnf config.Config) *Clients {
	conn, err := grpc.NewClient("black_list" + cnf.BLACKLIST_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("error:", err)
	}


	adminC := pb.NewAdminServiceClient(conn)
	hrC := pb.NewHRServiceClient(conn)
	blacklistC := pb.NewBlackListServiceClient(conn)

	return &Clients{
		AdminClient:     adminC,
		BlacklistClient:  blacklistC,
		HrClient:        hrC,
	}
}
