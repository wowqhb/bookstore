package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/timex"
	"github.com/wowqhb/bookstore/rpc/check/check"
	"github.com/wowqhb/bookstore/rpc/check/internal/svc"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {
	// todo: add your logic here and delete this line
	// 手动代码开始
	start := timex.Now()
	resp, err := l.svcCtx.Model.FindOne(in.Book)
	duration := timex.Since(start)
	if duration > time.Second {
		logx.WithDuration(duration).Slowf("[SQL] slow call")
	} else {
		logx.WithDuration(duration).Infof("[SQL] query")
	}
	if err != nil {
		return nil, err
	}

	return &check.CheckResp{
		Found: true,
		Price: resp.Price,
	}, nil
	// 手动代码结束
}
