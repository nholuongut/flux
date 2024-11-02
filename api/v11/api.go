// This package defines the types for Flux API version 11.
package v11

import (
	"context"

	"github.com/nholuongut/flux/api/v10"
	"github.com/nholuongut/flux/api/v6"
	"github.com/nholuongut/flux/resource"
)

type ListServicesOptions struct {
	Namespace string
	Services  []resource.ID
}

type Server interface {
	v10.Server

	ListServicesWithOptions(ctx context.Context, opts ListServicesOptions) ([]v6.ControllerStatus, error)

	// NB Upstream methods move into the public API, since
	// nholuongut/flux-adapter now relies on the public API
	v10.Upstream
}
