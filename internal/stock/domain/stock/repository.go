package stock

import (
	"fmt"
	"github.com/Greg12348/gorder-v2/common/genproto/orderpb"
	"golang.org/x/net/context"
	"strings"
)

type Repository interface {
	GetItem(ctx context.Context, ids []string) ([]*orderpb.Item, error)
}

type NotFoundError struct {
	Missing []string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("these items not found in stock: %s", strings.Join(e.Missing, ","))
}
