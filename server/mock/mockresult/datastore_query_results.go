// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import (
	"context"

	"github.com/fleetdm/fleet/v4/server/fleet"
)

var _ fleet.QueryResultStore = (*QueryResultStore)(nil)

type WriteResultFunc func(result fleet.DistributedQueryResult) error

type ReadChannelFunc func(ctx context.Context, query fleet.DistributedQueryCampaign) (<-chan interface{}, error)

type HealthCheckFunc func() error

type QueryResultStore struct {
	WriteResultFunc        WriteResultFunc
	WriteResultFuncInvoked bool

	ReadChannelFunc        ReadChannelFunc
	ReadChannelFuncInvoked bool

	HealthCheckFunc        HealthCheckFunc
	HealthCheckFuncInvoked bool
}

func (s *QueryResultStore) WriteResult(result fleet.DistributedQueryResult) error {
	s.WriteResultFuncInvoked = true
	return s.WriteResultFunc(result)
}

func (s *QueryResultStore) ReadChannel(ctx context.Context, query fleet.DistributedQueryCampaign) (<-chan interface{}, error) {
	s.ReadChannelFuncInvoked = true
	return s.ReadChannelFunc(ctx, query)
}

func (s *QueryResultStore) HealthCheck() error {
	s.HealthCheckFuncInvoked = true
	return s.HealthCheckFunc()
}