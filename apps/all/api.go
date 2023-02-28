package all

import (
    // 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
    _ "github.com/motongxue/cmdb-g7/apps/book/api"
    _ "github.com/motongxue/cmdb-g7/apps/host/api"
    _ "github.com/motongxue/cmdb-g7/apps/resource/api"
    _ "github.com/motongxue/cmdb-g7/apps/secret/api"
    _ "github.com/motongxue/cmdb-g7/apps/task/api"
)

// TODO 待添加各个模块的readme、以及mcenter为v0.0.5版本！
