package daemon

import (
	"github.com/nholuongut/flux/job"
	"github.com/nholuongut/flux/update"
)

type note struct {
	JobID  job.ID        `json:"jobID"`
	Spec   update.Spec   `json:"spec"`
	Result update.Result `json:"result"`
}
