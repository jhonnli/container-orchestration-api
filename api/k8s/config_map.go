package k8s

import (
	"github.com/jhonnli/container-orchestration-api/model/k8s"
	apiV1 "k8s.io/api/core/v1"
)

type ConfigMapInterface interface {
	Create(env, namespace string, configMap k8s.ConfigMapParam) (*apiV1.ConfigMap, error)
	List(env, namespace string) ([]apiV1.ConfigMap, error)
	Update(env, namespace string, configMap k8s.ConfigMapParam) (*apiV1.ConfigMap, error)
	Delete(env, namespace, configMapName string) error
	Get(env, namespace, configMapName string) (*apiV1.ConfigMap, error)
}
