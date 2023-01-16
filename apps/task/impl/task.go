package impl

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/motongxue/cmdb-g7/apps/host"
	"github.com/motongxue/cmdb-g7/apps/resource"
	"github.com/motongxue/cmdb-g7/apps/secret"
	"github.com/motongxue/cmdb-g7/apps/task"
	"github.com/motongxue/cmdb-g7/conf"
	"github.com/motongxue/cmdb-g7/provider/txyun/connectivity"
	"github.com/motongxue/cmdb-g7/provider/txyun/cvm"
)

func (i *impl) CreateTask(ctx context.Context, req *task.CreateTaskRequst) (*task.Task, error) {
	// 1. 查询secret
	s, err := i.secret.DescribeSecret(ctx, secret.NewDescribeSecretRequest(req.SecretId))
	if err != nil {
		return nil, err
	}
	// 解密api secret
	if err := s.Data.DecryptAPISecret(conf.C().App.EncryptKey); err != nil {
		return nil, err
	}
	// 分类参数类型
	switch req.Type {
	// 资源同步
	case task.Type_RESOURCE_SYNC:
		// 根据secret所属厂商，初始化对应厂商的operator
		switch s.Data.Vendor {
		case resource.Vendor_TENCENT:
			switch req.ResourceType {
			case resource.Type_HOST:
				// 实现主机资源同步，初始化腾讯cvm operator
				txConn := connectivity.NewTencentCloudClient(s.Data.ApiKey, s.Data.ApiSecret, req.Region)
				cvmOp := cvm.NewCVMOperator(txConn.CvmClient())

				// 同步所有资源时，进行分页查询
				pagger := cvm.NewPagger(float64(s.Data.RequestRate), cvmOp)
				for pagger.Next() {
					set := host.NewHostSet()
					// 查询分页有错误则返回
					if err := pagger.Scan(ctx, set); err != nil {
						return nil, err
					}
					// 保持该页数据、同步时间，记录到日志中
					for index := range set.Items {
						_, err := i.host.SyncHost(ctx, set.Items[index])
						if err != nil {
							i.log.Errorf("sync host error, %s", err)
							continue
						}
					}
				}
			case resource.Type_RDS:
			case resource.Type_BILL:
			}
		case resource.Vendor_ALIYUN:
		case resource.Vendor_HUAWEI:
		case resource.Vendor_AMAZON:
		case resource.Vendor_VSPHERE:
		default:
			return nil, fmt.Errorf("unknow resource type: %s", s.Data.Vendor)

		}
		// 资源释放
	case task.Type_RESOURCE_RELEASE:
	default:
		return nil, fmt.Errorf("unkown task type: %s", req.Type)
	}
	return nil, nil
}
func (i *impl) QueryBook(ctx context.Context, req *task.QueryTaskRequest) (*task.TaskSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBook not implemented")
}
func (i *impl) DescribeBook(ctx context.Context, req *task.DescribeTaskRequest) (*task.Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeBook not implemented")
}
