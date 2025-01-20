package repository

import (
	"context"

	"github.com/sk-pathak/fissionNode-backend/internal/db"
)

type NodeRepository struct {
	queries *db.Queries
}

func NewNodeRepository(queries *db.Queries) *NodeRepository {
	return &NodeRepository{queries: queries}
}

func (r *NodeRepository) RegisterNode(ctx context.Context, node *db.Node) error {
	_, err := r.queries.RegisterNode(ctx, db.RegisterNodeParams{
		NodeName:  node.NodeName,
		IpAddress: node.IpAddress,
		Status:    node.Status,
		Capacity:  node.Capacity,
	})

	return err
}

func (r *NodeRepository) GetNodeByID(ctx context.Context, id int64) (db.Node, error) {
	node, err := r.queries.GetNodeByID(ctx, id)
	if err != nil {
		return db.Node{}, err
	}

	result := db.Node{
		ID:            node.ID,
		IpAddress:     node.IpAddress,
		Capacity:      node.Capacity,
		Status:        node.Status,
		NodeName:      node.NodeName,
		LastHeartbeat: node.LastHeartbeat,
	}

	return result, nil
}

func (r *NodeRepository) GetNodes(ctx context.Context) ([]db.Node, error) {
	nodes, err := r.queries.GetNodes(ctx)
	if err != nil {
		return nil, err
	}

	var result []db.Node
	for _, node := range nodes {
		result = append(result, db.Node{
			ID:            node.ID,
			IpAddress:     node.IpAddress,
			Capacity:      node.Capacity,
			Status:        node.Status,
			NodeName:      node.NodeName,
			LastHeartbeat: node.LastHeartbeat,
		})
	}

	return result, nil
}

func (r *NodeRepository) DeleteNode(ctx context.Context, id int64) error {
	err := r.queries.DeleteNode(ctx, id)

	return err
}

func (r *NodeRepository) UpdateStatus(ctx context.Context, node *db.Node) error {
	err := r.queries.UpdateStatus(ctx, db.UpdateStatusParams{
		ID:            node.ID,
		Status:        node.Status,
		LastHeartbeat: node.LastHeartbeat,
	})

	return err
}
