package connectivity_test

import (
	"github.com/motongxue/cmdb-g7/provider/txyun/connectivity"
	"testing"
)

func TestTencentCloudClient(t *testing.T) {
	conn := connectivity.C()
	if err := conn.Check(); err != nil {
		t.Fatal(err)
	}
	t.Log(conn.AccountID())
}

func init() {
	//  初始化client
	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
