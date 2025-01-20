package service

import (
	"context"
	"errors"

	"github.com/sk-pathak/fissionNode-backend/internal/app/repository"
	"github.com/sk-pathak/fissionNode-backend/internal/db"
)

type NodeService struct {
	repo *repository.NodeRepository
}

func NewNodeService(repo *repository.NodeRepository) *NodeService {
	return &NodeService{repo: repo}
}

func (s *NodeService) RegisterNode(ctx context.Context, node *db.Node) error {
	err := s.repo.RegisterNode(ctx, node)
	if err != nil {
		return errors.New("Couldn't create node: " + err.Error())
	}
	return nil
}

func (s *NodeService) GetNodeByID(ctx context.Context, id int64) (db.Node, error) {
	node, err := s.repo.GetNodeByID(ctx, id)
	if err != nil {
		return db.Node{}, errors.New("Falied to fetch node: " + err.Error())
	}
	return node, nil
}

func (s *NodeService) GetNodes(ctx context.Context) ([]db.Node, error) {
	nodes, err := s.repo.GetNodes(ctx)
	if err != nil {
		return nil, errors.New("Failed to fetch nodes: " + err.Error())
	}
	return nodes, nil
}

func (s *NodeService) DeleteNode(ctx context.Context, id int64) error {
	_, err := s.repo.GetNodeByID(ctx, id)
	if err!=nil {
		return errors.New("Node does not exist: " + err.Error())
	}

	err = s.repo.DeleteNode(ctx, id)
	return err
}

func (s *NodeService) UpdateStatus(ctx context.Context, node* db.Node) error {
	err := s.repo.UpdateStatus(ctx, node)
	return err
}
