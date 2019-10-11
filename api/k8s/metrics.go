package k8s

import (
	"github.com/jhonnli/container-orchestration-api/model/metrics"
)

type MetricsInterface interface {
	ListMetricsOfNode(env string) ([]metrics.NodeMetrics, error)

	GetMetricsOfNodeByNodeName(env, node string) (*metrics.NodeMetrics, error)

	ListMetricsOfNamespace(env, namespace string) ([]metrics.PodMetrics, error)

	GetMetricsOfNamespaceByPodName(env, namespace, podName string) (*metrics.PodMetrics, error)
}
