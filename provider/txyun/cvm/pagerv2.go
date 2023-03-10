package cvm

import (
	"context"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/motongxue/cmdb-g7/apps/host"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	tx_cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func NewPagerV2(op *CVMOperator) host.PagerV2 {
	return &pagerV2{
		op:          op,
		log:         zap.L().Named("CVM"),
		BasePagerV2: host.NewBasePagerV2(),
		req:         tx_cvm.NewDescribeInstancesRequest(),
	}
}

type pagerV2 struct {
	req *cvm.DescribeInstancesRequest
	op  *CVMOperator
	log logger.Logger
	*host.BasePagerV2
}

// 修改Req 执行真正的下一页的offset
func (p *pagerV2) nextReq() *cvm.DescribeInstancesRequest {
	os := p.Offset()
	ps := p.PageSize()
	p.req.Offset = &os
	p.req.Limit = &ps
	return p.req
}

func (p *pagerV2) Scan(ctx context.Context, set host.Set) error {
	p.log.Debugf("query page: %d", p.PageNumber())
	hs, err := p.op.Query(ctx, p.nextReq())
	if err != nil {
		return err
	}

	// 把查询出来的数据赋值给set
	for i := range hs.Items {
		set.Add(hs.Items[i])
	}

	// 可以根据当前一页是满页来决定是否有下一页
	p.CheckHasNext(hs.Length())
	return nil
}
