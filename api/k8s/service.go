package k8s

import (
	"github.com/jhonnli/container-orchestration-api/model/k8s"
	v12 "k8s.io/api/core/v1"
)

type ServiceInterface interface {
	Get(env, namespace string, svcname string) (*v12.Service, error)
	Delete(env, namespace string, svcname string) error
	List(env, namespace string) (*v12.ServiceList, error)
	Create(env string, param k8s.ServiceParam) (*v12.Service, error)
	Update(env string, param k8s.ServiceParam) (*v12.Service, error)
	Apply(env string, param k8s.ServiceParam) (*v12.Service, error)
}
