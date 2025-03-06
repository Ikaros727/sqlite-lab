package v1

import (
	"context"

	"github.com/Ikaros727/sqlite-lab/internal/svc"
	"github.com/Ikaros727/sqlite-lab/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ExecLogic
type ExecLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExecLogic {
	return &ExecLogic{
		Logger: logx.WithContext(ctx).WithFields(logx.Field("Logic", "ExecLogic")),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExecLogic) Exec(req *types.ExecReq) (resp *types.ExecResp, err error) {
	resp = &types.ExecResp{}
	resp.Status = types.StatusSuccess()

	// 获取数据库操作对象
	db, err := l.svcCtx.UserDatabaseCache.MustLoad(1, req.Database)
	if err != nil {
		l.Logger.Errorf("UserDatabaseCache.MustLoad(%v, %v) failed, err: %v", 1, req.Database, err)
		return
	}

	// 参数转换
	args := make([]interface{}, len(req.Args))
	for i, arg := range req.Args {
		args[i] = arg
	}

	// 执行操作
	db = db.Exec(req.SQL, args...)
	if err = db.Error; err != nil {
		l.Logger.Errorf("db.Exec failed, err: %v", err)
		return
	}
	resp.Data.RowsAffected = db.RowsAffected
	if err = db.Raw("SELECT last_insert_rowid()").Scan(&resp.Data.ID).Error; err != nil {
		l.Logger.Errorf("db.Raw.Scan failed, err: %v", err)
		return
	}
	return
}
