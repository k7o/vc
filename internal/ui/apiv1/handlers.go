package apiv1

import (
	"context"
	"errors"
	"time"
	apiv1_apigw "vc/internal/apigw/apiv1"
	"vc/internal/gen/status/apiv1_status"
	"vc/pkg/model"
)

func (c *Client) Status(ctx context.Context, req *apiv1_status.StatusRequest) (*apiv1_status.StatusReply, error) {
	probes := model.Probes{}
	status := probes.Check("ui")
	return status, nil
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoggedinReply struct {
	Username string `json:"username" binding:"required"`
	// LoggedInTime RFC3339
	LoggedInTime time.Time `json:"logged_in_time" binding:"required"`
}

func (c *Client) Login(ctx context.Context, req *LoginRequest) (*LoggedinReply, error) {
	if req.Username != c.cfg.UI.Username || req.Password != c.cfg.UI.Password {
		return nil, errors.New("invalid username and/or password")
	}

	reply := &LoggedinReply{
		Username:     c.cfg.UI.Username,
		LoggedInTime: time.Now(),
	}

	return reply, nil
}

func (c *Client) Logout(ctx context.Context) error {
	return nil
}

func (c *Client) User(ctx context.Context) (*LoggedinReply, error) {
	return nil, nil
}

type DocumentListRequest struct {
	AuthenticSource string          `json:"authentic_source"`
	Identity        *model.Identity `json:"identity" validate:"required"`
	DocumentType    string          `json:"document_type"`
	ValidFrom       int64           `json:"valid_from"`
	ValidTo         int64           `json:"valid_to"`
}

type PortalRequest struct {
	DocumentType            string `json:"document_type" binding:"required"`
	AuthenticSource         string `json:"authentic_source" binding:"required"`
	AuthenticSourcePersonId string `json:"authentic_source_person_id" binding:"required"`
}

func (c *Client) DocumentList(ctx context.Context, req *DocumentListRequest) (any, error) {
	reply, err := c.apigwClient.DocumentList(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (c *Client) Upload(ctx context.Context, req *apiv1_apigw.UploadRequest) (any, error) {
	reply, err := c.apigwClient.Upload(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

type MockNextRequest struct {
	PortalRequest
}

func (c *Client) MockNext(ctx context.Context, req *MockNextRequest) (any, error) {
	if c.cfg.Common.Kafka.Enabled {
		if err := c.eventPublisher.MockNext(req); err != nil {
			return nil, err
		}
		return nil, nil
	}

	reply, err := c.mockasClient.MockNext(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (c *Client) StatusAPIGW(ctx context.Context, req *apiv1_status.StatusRequest) (any, error) {
	reply, err := c.apigwClient.Status()
	if err != nil {
		return nil, err
	}
	return reply, nil
}
