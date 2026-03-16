package api

import (
	"context"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) ListNamespaces(ctx context.Context, request ListNamespacesRequestObject) (ListNamespacesResponseObject, error) {
	return ListNamespaces200JSONResponse{}, nil
}

func (Server) CreateNamespace(ctx context.Context, request CreateNamespaceRequestObject) (CreateNamespaceResponseObject, error) {
	return CreateNamespace201JSONResponse{}, nil
}

func (Server) DeleteNamespace(ctx context.Context, request DeleteNamespaceRequestObject) (DeleteNamespaceResponseObject, error) {
	return DeleteNamespace204Response{}, nil
}

func (Server) GetNamespace(ctx context.Context, request GetNamespaceRequestObject) (GetNamespaceResponseObject, error) {
	return GetNamespace200JSONResponse{}, nil
}

func (Server) UpdateNamespace(ctx context.Context, request UpdateNamespaceRequestObject) (UpdateNamespaceResponseObject, error) {
	return UpdateNamespace200JSONResponse{}, nil
}
