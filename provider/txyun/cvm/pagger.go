package cvm

import (
	"context"
	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/motongxue/cmdb-g7/apps/host"
	"time"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	tx_cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func NewPagger(op *CVMOperator) host.Pagger {
	p := &pagger{
		op:         op,
		hasNext:    true,
		pageNumber: 1,
		pageSize:   20,
		log:        zap.L().Named("CVM"),
		tb:         tokenbucket.NewBucket(1*time.Second, 3),
	}

	p.req = tx_cvm.NewDescribeInstancesRequest()
	p.req.Limit = &p.pageSize
	p.req.Offset = p.offset()
	return p
}

type pagger struct {
	req     *cvm.DescribeInstancesRequest
	op      *CVMOperator
	hasNext bool
	// 令牌桶
	tb *tokenbucket.Bucket

	// 控制分页的核心参数
	pageNumber int64
	pageSize   int64
	log        logger.Logger
}

func (p *pagger) SetPageSize(ps int64) {
	p.pageSize = ps
}

// 根据分页参数来计算
func (p *pagger) offset() *int64 {
	offSet := (p.pageNumber - 1) * p.pageSize
	return &offSet
}

// 需要在请求数据是 计算出来(根据当前页是否满)
func (p *pagger) Next() bool {
	return p.hasNext
}

// 修改Req 执行真正的下一页的offset
func (p *pagger) nextReq() *cvm.DescribeInstancesRequest {
	// 等待分配令牌
	p.tb.Wait(1)

	p.req.Offset = p.offset()
	p.req.Limit = &p.pageSize
	return p.req
}

func (p *pagger) Scan(ctx context.Context, set *host.HostSet) error {
	p.log.Debugf("query page: %d", p.pageNumber)
	hs, err := p.op.Query(ctx, p.nextReq())
	if err != nil {
		return err
	}
	// 把查询出来的数据赋值给set
	*set = *hs.Clone()

	// 可以根据当前一页是满页来决定是否有下一页
	if hs.Length() < p.pageSize {
		p.hasNext = false
	}

	// 直接调整指针到下一页
	p.pageNumber++
	return nil
}
