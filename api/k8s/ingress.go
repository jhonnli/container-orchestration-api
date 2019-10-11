package k8s

import (
	"github.com/jhonnli/container-orchestration-api/model/k8s"
	v1beta12 "k8s.io/api/extensions/v1beta1"
)

type IngressInterface interface {
	Apply(env string, param k8s.IngressParam) (*v1beta12.Ingress, error)
	Update(env string, param k8s.IngressParam) (*v1beta12.Ingress, error)

	Get(env, namespace, ingName string) (*v1beta12.Ingress, error)

	List(env, namespace string) (*v1beta12.IngressList, error)

	Create(env string, param k8s.IngressParam) (*v1beta12.Ingress, error)

	Delete(env, namespace, ingName string) error
}
