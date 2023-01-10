package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/motongxue/cmdb-g7/apps/host"
	"github.com/motongxue/cmdb-g7/conf"
)

var (
	svr = &service{}
)

type service struct {
	db *sql.DB

	log logger.Logger
	host.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	s.db = db
	if err != nil {
		return err
	}
	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *service) Name() string {
	return host.AppName
}
func (s *service) Registry(server *grpc.Server) {
	host.RegisterServiceServer(server, s)
}
func init() {
	app.RegistryGrpcApp(svr)
}
