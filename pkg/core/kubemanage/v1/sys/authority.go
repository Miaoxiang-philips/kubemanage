package sys

import (
	"context"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
)

type AuthorityGetter interface {
	Authority() Authority
}

type Authority interface {
	SetMenuAuthority(ctx context.Context, auth *model.SysAuthority) error
}

type authority struct {
	factory dao.ShareDaoFactory
}

func NewAuthority(factory dao.ShareDaoFactory) *authority {
	return &authority{factory: factory}
}

func (a *authority) SetMenuAuthority(ctx context.Context, auth *model.SysAuthority) error {
	return a.factory.Authority().SetMenuAuthority(ctx, auth)
}