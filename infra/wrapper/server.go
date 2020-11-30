package wrapper

import (
	"context"
	err2 "gimServer/infra/err"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewServerWrapper() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if e := recover(); e != nil {
				logrus.Errorln(ctx, "发生panic", e)
			}
		}()
		//deadline, _ := ctx.Deadline()
		////如果超时5s在deadline之后，则重置deadline为5s后
		//if time.Now().Add(time.Second * 5).After(deadline) {
		//	ctx, _ = context.WithTimeout(ctx, time.Second*5)
		//}
		resp, err = handler(ctx, req)
		logrus.Infoln(req, err)
		if err != nil {
			evaError := err2.FromError(err)
			return resp, err2.EncodeStatus(evaError).Err()
		}
		return resp, err
	}
}
