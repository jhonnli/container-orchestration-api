package k8s

import (
	"github.com/jhonnli/container-orchestration-api/model/k8s"
	v12 "k8s.io/api/core/v1"
)

type NamespaceInterface interface {
	Create(env string, param k8s.NamespaceParam) (*v12.Namespace, error)

	List(env string) (*v12.NamespaceList, error)

	Delete(env, namespace string) error

	Get(env, namespace string) (*v12.Namespace, error)
}
