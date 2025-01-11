package grpc

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/application"
	"github.com/kidus-tiliksew/aqua-crims/grpc/proto"
)

type NotificationGRPCServer struct {
	proto.UnimplementedNotificationServiceServer
	app application.App
}

func NewNotificationGRPCServer(app application.App) *NotificationGRPCServer {
	return &NotificationGRPCServer{app: app}
}

func (s *NotificationGRPCServer) DeleteNotification(ctx context.Context, req *proto.DeleteNotificationRequest) (*proto.DeleteNotificationReply, error) {
	if err := s.app.NotificationDelete(ctx, req.Id); err != nil {
		return nil, err
	}

	return &proto.DeleteNotificationReply{Message: "Notification deleted"}, nil
}

func (s *NotificationGRPCServer) DeleteNotificationByUser(ctx context.Context, req *proto.DeleteNotificationByUserRequest) (*proto.DeleteNotificationByUserReply, error) {
	if err := s.app.NotificationDeleteByUser(ctx, req.UserId); err != nil {
		return nil, err
	}

	return &proto.DeleteNotificationByUserReply{Message: "All notifications deleted for user"}, nil
}

func (s *NotificationGRPCServer) GetNotificationByUser(ctx context.Context, req *proto.GetNotificationByUserRequest) (*proto.GetNotificationByUserReply, error) {
	notifications, err := s.app.NotificationGetByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var protoNotifications []*proto.Notification
	for _, n := range notifications {
		protoNotifications = append(protoNotifications, &proto.Notification{
			Id:      int32(n.ID),
			UserId:  n.UserID,
			Message: n.Message,
		})
	}

	return &proto.GetNotificationByUserReply{Notifications: protoNotifications}, nil
}
