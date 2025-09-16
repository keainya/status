package main

import (
	"context"

	keainya "github.com/keainya/status/kitex_gen/keainya"
)

// StatusServiceImpl implements the last service interface defined in the IDL.
type StatusServiceImpl struct{}

// Status implements the StatusServiceImpl interface.
func (s *StatusServiceImpl) Status(ctx context.Context) (resp *keainya.BaseResp, err error) {
	return &keainya.BaseResp{
		Code: 0,
		Msg:  "keainya central server",
	}, nil
}
