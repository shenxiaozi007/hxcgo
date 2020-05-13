package uuid

import "github.com/bwmarrin/snowflake"

type IDFactory struct {
	node *snowflake.Node
}

func New(n int64) (*IDFactory, error) {
	node, err := snowflake.NewNode(n)
	if err != nil {
		return nil, err
	}

	return &IDFactory{node: node}, nil
}

func (id *IDFactory) GetID() int64 {
	return id.node.Generate().Int64()
}

func (id *IDFactory) String() string {
	return id.node.Generate().String()
}
