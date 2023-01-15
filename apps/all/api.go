package all

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/motongxue/cmdb-g7/apps/book/api"
	_ "github.com/motongxue/cmdb-g7/apps/host/api"
	_ "github.com/motongxue/cmdb-g7/apps/resource/api"
	_ "github.com/motongxue/cmdb-g7/apps/secret/api"
)
