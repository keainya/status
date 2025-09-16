package main

import (
	"context"

	status "github.com/keainya/status/kitex_gen/status"
)

// StatusServiceImpl implements the last service interface defined in the IDL.
type StatusServiceImpl struct{}

// Status implements the StatusServiceImpl interface.
func (s *StatusServiceImpl) Status(ctx context.Context) (resp *status.StatusInfo, err error) {
	return &status.StatusInfo{
		Code: 0,
		Msg:  "keainya central server!",
	}, nil
}
