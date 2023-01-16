package impl_test

import (
	"context"
	"os"
	"testing"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/motongxue/cmdb-g7/apps/resource"
	"github.com/motongxue/cmdb-g7/apps/task"
	"github.com/motongxue/cmdb-g7/conf"
)

var (
	ins task.ServiceServer
)

func TestImpl_CreateTask(t *testing.T) {
	req := task.NewCreateTaskRequst()
	req.Type = task.Type_RESOURCE_SYNC
	req.Region = os.Getenv("TX_CLOUD_REGION")
	req.ResourceType = resource.Type_HOST
	req.SecretId = "ca8sio13n7plv10h6fu0"
	taskIns, err := ins.CreateTask(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(taskIns)
}

func init() {
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
	zap.DevelopmentSetup()
	err := app.InitAllApp()
	if err != nil {
		panic(err)
	}
	ins = app.GetGrpcApp(task.AppName).(task.ServiceServer)
}
