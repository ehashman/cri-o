package v1alpha2

import (
	"context"

	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

func (s *service) Version(
	ctx context.Context, req *pb.VersionRequest,
) (*pb.VersionResponse, error) {
	resp, err := s.server.Version(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.VersionResponse{
		Version:           resp.Version,
		RuntimeName:       resp.RuntimeName,
		RuntimeVersion:    resp.RuntimeVersion,
		RuntimeApiVersion: resp.RuntimeAPIVersion,
	}, nil
}
