package k8s

import (
	coreV1 "k8s.io/api/core/v1"
)

type NodeInterface interface {
	List(env string) (*coreV1.NodeList, error)
	Get(env, nodeName string) (*coreV1.Node, error)
	Drain(cluster, nodeName string) error
	Cordon(cluster, nodeName string) error
	Uncordon(cluster, nodeName string) error
	AddNodeLabels(env, nodeName string, nodeLabels map[string]string) error
	DeleteNodeLabels(env, nodeName string, nodeLabels map[string]string) error
	ListNodeLabels(env, labelSelector string) (*coreV1.NodeList, error)
}
