package k8s

import (
	"github.com/jhonnli/container-orchestration-api/model/k8s"
	coreV1 "k8s.io/api/core/v1"
)

type PodInterface interface {
	ListByNamespace(envName, namespace string) (*coreV1.PodList, error)
	ListPodByDeployment(env, nsname, deployName string) (*coreV1.PodList, error)
	DeletePodByDeployment(env, nsname, deployName string) error
	//ListByNamespaceAndOptions(envName, namespace string, options k8s.ListOptions) (*coreV1.PodList, error)
	ListByNode(env, nodeName string) (*coreV1.PodList, error)
	GetByName(env, namespace, podName string) (*coreV1.Pod, error)
	Eviction(env, namespace, podName string) error
	EvictionPods(env, namespace string, data []string) (bool, error)
	HealthCheck(env, namespace, podPrefix string) (k8s.PodHealthResponse, error)
	Log(env, namespace, podName string, logOptionParam k8s.PodLogOptionParam) (string, error)
}
