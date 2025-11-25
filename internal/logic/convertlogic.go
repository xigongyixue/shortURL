// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"shortURL/internal/svc"
	"shortURL/internal/types"
	"shortURL/pkg/connect"
	"shortURL/pkg/md5"
	"shortURL/pkg/urltool"
	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 转链：长链接-> 短链接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 1. 校验长链接
	// 1.1 非空
	// 1.2 有效
	if ok := connect.Get(req.LongUrl); !ok {
		return nil, errors.New("无效的长链接")
	}
	// 1.3 未被转换过
	md5Value := md5.Sum([]byte(req.LongUrl))
	u, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{String: md5Value, Valid: true})
	if err != sql.ErrNoRows {
		if err == nil {
			return nil, fmt.Errorf("长链接已被转换为%s", u.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", logx.LogField{Key:"err", Value: err})
		return nil, err
	}
	// 1.4 不是被使用过的短链接
	basePath, err := urltool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("urltool.GetBasePath failed", logx.LogField{Key:"lurl", Value: req.LongUrl})
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: basePath, Valid: true})
	if err != sql.ErrNoRows {
		if err == nil {
			return nil, errors.New("长链接已被转换为短链接")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{Key:"err", Value: err})
		return nil, err
	}
	// 2. 取号
	// 每来一个转链请求，就是用replace into语句往sequence表插入一条数据，并取出主键id作为号码
	seq, err := l.svcCtx.Sequence.Next()
	if err != nil {
		logx.Errorw("Sequence.Next failed", logx.LogField{Key:"err", Value: err})
		return nil, err
	}
	fmt.Println(seq)
	// 3. 号码转短链
	// 4. 存储映射关系
	// 5. 返回结果

	return
}
