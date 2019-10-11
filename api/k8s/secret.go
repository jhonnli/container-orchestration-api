package k8s

import (
	"github.com/jhonnli/container-orchestration-api/model/k8s"
	apiV1 "k8s.io/api/core/v1"
)

type SecretInterface interface {
	Create(env, namespace string, secretParam k8s.SecretParam) (*apiV1.Secret, error)

	List(env, namespace string) ([]apiV1.Secret, error)

	Update(env, namespace string, secret k8s.SecretParam) (*apiV1.Secret, error)

	Delete(env, namespace, secretName string) error

	Get(env, namespace, secretName string) (*apiV1.Secret, error)
}
