package main

import (
	"encoding/json"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/iochti/thing-group-service/models"
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
	if in.ID == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Error: bad ID")
	}
	group := new(models.ThingGroup)
	if err := t.Db.GetGroupByID(in.GetID(), group); err != nil {
		return nil, grpc.Errorf(codes.NotFound, err.Error())
	}
	bytes, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}
	return &pb.ThingGroup{Item: bytes}, nil
}

// ListUserGroups lists all groups for a given user
func (t *ThingGroupSvc) ListUserGroups(in *pb.UserIDRequest, stream pb.ThingGroupSvc_ListUserGroupsServer) error {
	fetched, err := t.Db.ListGroupsByUser(in.GetUserId())
	if err != nil {
		return grpc.Errorf(codes.Internal, err.Error())
	}
	for _, v := range fetched {
		bytes, err := json.Marshal(v)
		if err != nil {
			return grpc.Errorf(codes.Internal, err.Error())
		}
		stream.Send(&pb.ThingGroup{Item: bytes})
	}
	return nil
}

// CreateGroup creates a group in the database in returns it
func (t *ThingGroupSvc) CreateGroup(ctx context.Context, in *pb.ThingGroup) (*pb.ThingGroup, error) {
	group := new(models.ThingGroup)
	if err := json.Unmarshal(in.GetItem(), &group); err != nil {
		return nil, err
	}

	if err := t.Db.CreateGroup(group); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	bytes, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}
	return &pb.ThingGroup{Item: bytes}, nil
}

// UpdateGroup updates a group with group passed as parameter
func (t *ThingGroupSvc) UpdateGroup(ctx context.Context, in *pb.ThingGroup) (*pb.ThingGroup, error) {
	group := new(models.ThingGroup)
	if err := json.Unmarshal(in.GetItem(), &group); err != nil {
		return nil, err
	}

	if err := t.Db.UpdateGroup(group); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	bytes, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}
	return &pb.ThingGroup{Item: bytes}, nil
}

// DeleteGroup deletes a group in the db identified by its ID
func (t *ThingGroupSvc) DeleteGroup(ctx context.Context, in *pb.GroupIDRequest) (*empty.Empty, error) {
	if err := t.Db.DeleteGroup(in.GetID()); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}
