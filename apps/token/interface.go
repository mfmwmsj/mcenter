package token

import context "context"

type Service interface {
	// 颁发Token
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// 撤销Token
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)
	// 切换Token空间
	ChangeNamespace(context.Context, *ChangeNamespaceRequest) (*Token, error)
	// RPC
	RPCServer
}
