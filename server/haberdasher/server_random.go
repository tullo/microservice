package haberdasher

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/pkg/errors"
	"github.com/tullo/microservice/rpc/haberdasher"
	"github.com/twitchtv/twirp"
)

// MakeHat produces a hat of mysterious, randomly-selected color!
func (svc *Server) MakeHat(ctx context.Context, size *haberdasher.Size) (*haberdasher.Hat, error) {
	if size.Centimeters <= 0 {
		return nil, twirp.InvalidArgumentError("Centimeters", "I can't make a hat that small!")
	}
	ci, err := randomInt(int64(len(color)))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ni, err := randomInt(int64(len(name)))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	hat := Hat{
		Size:  uint32(size.Centimeters),
		Color: color[ci],
		Name:  name[ni],
	}
	hat.ID, err = svc.sonyflake.NextID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate sonyflake id: %w", err)
	}

	fields := strings.Join(HatFields, ",")
	named := ":" + strings.Join(HatFields, ",:")

	query := fmt.Sprintf("insert into %s (%s) values (%s)", HatTable, fields, named)
	_, err = svc.db.NamedExecContext(ctx, query, &hat)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error inserting hat: %s", query))
	}

	return &haberdasher.Hat{
		Size:  hat.Size,
		Color: hat.Color,
		Name:  hat.Name,
	}, nil
}

func randomInt(max int64) (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return -1, fmt.Errorf("failed to generate random int: %w", err)
	}

	return n.Int64(), nil
}
