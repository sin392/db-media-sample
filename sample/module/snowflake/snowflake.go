package snowflake

import (
	"context"
	"fmt"

	"github.com/bwmarrin/snowflake"
)

type SnowflakeIDKey struct{}

type SnowflakeIDGenerator struct {
	node *snowflake.Node
}

func NewSnowflakeIDGenerator(nodeID int64) (*SnowflakeIDGenerator, error) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to create snowflake node: %w", err)
	}
	return &SnowflakeIDGenerator{
		node: node,
	}, nil
}

func (g *SnowflakeIDGenerator) Generate() string {
	return g.node.Generate().String()
}

func SetSnowflakeID(ctx context.Context, snowflakeID string) context.Context {
	return context.WithValue(ctx, SnowflakeIDKey{}, snowflakeID)
}

func GetSnowflakeID(ctx context.Context) string {
	snowflakeID, _ := ctx.Value(SnowflakeIDKey{}).(string)
	return snowflakeID
}
