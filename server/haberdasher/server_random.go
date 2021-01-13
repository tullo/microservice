package haberdasher

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/tullo/microservice/rpc/haberdasher"
	"github.com/twitchtv/twirp"
)

// MakeHat produces a hat of mysterious, randomly-selected color!
func (svc *Server) MakeHat(ctx context.Context, size *haberdasher.Size) (*haberdasher.Hat, error) {
	if size.Centimeters <= 0 {
		return nil, twirp.InvalidArgumentError("Centimeters", "I can't make a hat that small!")
	}
	ci, err := randomInt(5)
	if err != nil {
		panic(err)
	}

	ni, err := randomInt(4)
	if err != nil {
		panic(err)
	}

	return &haberdasher.Hat{
		Size:  size.Centimeters,
		Color: []string{"white", "black", "brown", "red", "blue"}[ci],
		Name:  []string{"bowler", "baseball cap", "top hat", "derby"}[ni],
	}, nil
}

func randomInt(max int64) (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return -1, fmt.Errorf("failed to generate random int: %w", err)
	}

	return n.Int64(), nil
}
