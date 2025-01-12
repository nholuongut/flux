package resource

import (
	"github.com/nholuongut/flux/image"
	"github.com/nholuongut/flux/resource"
)

type Deployment struct {
	baseObject
	Spec DeploymentSpec
}

type DeploymentSpec struct {
	Replicas int
	Template PodTemplate
}

func (d Deployment) Containers() []resource.Container {
	return d.Spec.Template.Containers()
}

func (d Deployment) SetContainerImage(container string, ref image.Ref) error {
	return d.Spec.Template.SetContainerImage(container, ref)
}

var _ resource.Workload = Deployment{}
