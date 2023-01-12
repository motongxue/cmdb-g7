package cvm

import (
	"github.com/motongxue/cmdb-g7/apps/host"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type pagger struct {
	req     *cvm.DescribeInstancesRequest
	op      *CVMOperator
	hasNext bool
}

func (p *pagger) Next() bool {
	return p.hasNext
}
func (p *pagger) Scan(set *host.HostSet) error {
	return nil
}
