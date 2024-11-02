package v6

import (
	"context"

	"github.com/nholuongut/flux/cluster"
	"github.com/nholuongut/flux/git"
	"github.com/nholuongut/flux/job"
	"github.com/nholuongut/flux/resource"
	"github.com/nholuongut/flux/ssh"
	"github.com/nholuongut/flux/update"
)

type ImageStatus struct {
	ID         resource.ID
	Containers []Container
}

// ReadOnlyReason enumerates the reasons that a controller is
// considered read-only. The zero value is considered "OK", since the
// zero value is what prior versions of the daemon will effectively
// send.
type ReadOnlyReason string

const (
	ReadOnlyOK       ReadOnlyReason = ""
	ReadOnlyMissing  ReadOnlyReason = "NotInRepo"
	ReadOnlySystem   ReadOnlyReason = "System"
	ReadOnlyNoRepo   ReadOnlyReason = "NoRepo"
	ReadOnlyNotReady ReadOnlyReason = "NotReady"
	ReadOnlyROMode   ReadOnlyReason = "ReadOnlyMode"
)

type ControllerStatus struct {
	ID         resource.ID
	Containers []Container
	ReadOnly   ReadOnlyReason
	Status     string
	Rollout    cluster.RolloutStatus
	SyncError  string
	Antecedent resource.ID
	Labels     map[string]string
	Automated  bool
	Locked     bool
	Ignore     bool
	Policies   map[string]string
}

// --- config types

type GitRemoteConfig struct {
	URL    string `json:"url"`
	Branch string `json:"branch"`
	Path   string `json:"path"`
}

type GitConfig struct {
	Remote       GitRemoteConfig   `json:"remote"`
	PublicSSHKey ssh.PublicKey     `json:"publicSSHKey"`
	Status       git.GitRepoStatus `json:"status"`
}

type Deprecated interface {
	SyncNotify(context.Context) error
}

type NotDeprecated interface {
	// from v5
	Export(context.Context) ([]byte, error)

	// v6
	ListServices(ctx context.Context, namespace string) ([]ControllerStatus, error)
	ListImages(ctx context.Context, spec update.ResourceSpec) ([]ImageStatus, error)
	UpdateManifests(context.Context, update.Spec) (job.ID, error)
	SyncStatus(ctx context.Context, ref string) ([]string, error)
	JobStatus(context.Context, job.ID) (job.Status, error)
	GitRepoConfig(ctx context.Context, regenerate bool) (GitConfig, error)
}

type Upstream interface {
	// from v4
	Ping(context.Context) error
	Version(context.Context) (string, error)
}

type Server interface {
	Deprecated
	NotDeprecated
}
