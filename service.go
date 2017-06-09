package main

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/iochti/thing-group-service/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// ThingGroupSvc implements ThingGroupSvc from protobuf
type ThingGroupSvc struct {
	Db DataLayerInterface
}

// GetGroup gets a thing-group from the database
func (t *ThingGroupSvc) GetGroup(ctx context.Context, in *pb.GroupIDRequest) (*pb.ThingGroup, error) {
	if in.ID <= 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "Error: bad ID")
	}
	group := new(pb.ThingGroup)
	if err := t.Db.GetGroupById(in.GetID(), group); err != nil {
		return nil, grpc.Errorf(codes.NotFound, err.Error())
	}
	return group, nil
}

// CreateGroup creates a group in the database in returns it
func (t *ThingGroupSvc) CreateGroup(ctx context.Context, in *pb.ThingGroup) (*pb.ThingGroup, error) {
	created := new(pb.ThingGroup)
	if err := t.Db.CreateGroup(created); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return created, nil
}

// UpdateGroup updates a group with group passed as parameter
func (t *ThingGroupSvc) UpdateGroup(ctx context.Context, in *pb.ThingGroup) (*pb.ThingGroup, error) {
	if err := t.Db.UpdateGroup(in); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return in, nil
}

// DeleteGroup deletes a group in the db identified by its ID
func (t *ThingGroupSvc) DeleteGroup(ctx context.Context, in *pb.GroupIDRequest) (*empty.Empty, error) {
	if err := t.Db.DeleteGroup(in.GetID()); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}
