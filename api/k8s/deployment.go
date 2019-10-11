package k8s

import (
	"github.com/jhonnli/container-orchestration-api/model/k8s"
	v12 "k8s.io/api/apps/v1"
)

type DeploymentInterface interface {
	Apply(env string, param k8s.DeploymentParam) (*v12.Deployment, error)
	Create(env string, param k8s.DeploymentParam) (*v12.Deployment, error)

	Update(env string, param k8s.DeploymentParam) (*v12.Deployment, error)

	Get(env, namespaceName, deployName string) (*v12.Deployment, error)
	List(env, namespaceName string) (*v12.DeploymentList, error)
	Delete(env, namespaceName, deployName string) error
}
