package authzed

import (
	"context"

	api "github.com/authzed/authzed-go/arrakisapi/api"
)

type WriteOp int

const (
	CREATE WriteOp = 1
	TOUCH  WriteOp = 2
	DELETE WriteOp = 3
)

type Client struct {
	grpcClient api.ACLServiceClient
}

func NewClient(token string) (*Client, error) {
	opts, err := NewGrpcClientOptions(token)
	if err != nil {
		return nil, err
	}

	client, err := NewGrpcClient(opts)
	if err != nil {
		return nil, err
	}

	return &Client{grpcClient: client}, nil
}

func (c *Client) Write(op WriteOp, tuples ...Tuple) (revision string, err error) {
	var tupleUpdates []*api.RelationTupleUpdate
	for _, tuple := range tuples {
		tupleUpdates = append(tupleUpdates, &api.RelationTupleUpdate{
			Operation: api.RelationTupleUpdate_Operation(op),
			Tuple:     tuple.asProto(),
		})
	}

	resp, err := c.grpcClient.Write(context.TODO(), &api.WriteRequest{Updates: tupleUpdates})
	if err != nil {
		return "", err
	}

	return resp.Revision.Token, nil
}

func (c *Client) Check(t Tuple, revision string) (bool, error) {
	protoTuple := t.asProto()
	resp, err := c.grpcClient.Check(context.TODO(), &api.CheckRequest{
		TestUserset: protoTuple.ObjectAndRelation,
		User:        protoTuple.User,
		AtRevision:  &api.Zookie{Token: revision},
	})

	return resp.IsMember, err
}

func (c *Client) ContentChangeCheck(t Tuple) (revision string, allowed bool, err error) {
	protoTuple := t.asProto()
	resp, err := c.grpcClient.ContentChangeCheck(context.TODO(), &api.ContentChangeCheckRequest{
		TestUserset: protoTuple.ObjectAndRelation,
		User:        protoTuple.User,
	})

	return resp.Revision.Token, resp.IsMember, err
}
