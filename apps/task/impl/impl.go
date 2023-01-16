package impl

import (
	"database/sql"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/motongxue/cmdb-g7/apps/host"
	"github.com/motongxue/cmdb-g7/apps/secret"
	"github.com/motongxue/cmdb-g7/apps/task"
	"github.com/motongxue/cmdb-g7/conf"
	"google.golang.org/grpc"
)

var (
	// service 服务实例
	svr = &impl{}
)

type impl struct {
	db  *sql.DB
	log logger.Logger
	task.UnimplementedServiceServer
	secret secret.ServiceServer
	host   host.ServiceServer
}

func (s *impl) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db
	s.log = zap.L().Named(s.Name())
	// 通过mock 来解耦以来 s.secret = &secretMock{}
	s.secret = app.GetGrpcApp(secret.AppName).(secret.ServiceServer)
	s.host = app.GetGrpcApp(host.AppName).(host.ServiceServer)
	return nil
}

func (s *impl) Name() string {
	return task.AppName
}

func (s *impl) Registry(server *grpc.Server) {
	task.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
