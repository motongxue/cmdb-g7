package all

import (
	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	_ "github.com/motongxue/cmdb-g7/apps/book/impl"
	_ "github.com/motongxue/cmdb-g7/apps/host/impl"
	_ "github.com/motongxue/cmdb-g7/apps/resource/impl"
	_ "github.com/motongxue/cmdb-g7/apps/secret/impl"
	_ "github.com/motongxue/cmdb-g7/apps/task/impl"
)
