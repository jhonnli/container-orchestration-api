package metrics

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/metrics/pkg/apis/metrics"
)

type NodeMetrics struct {
	metav1.TypeMeta
	metav1.ObjectMeta `json:"metadata"`

	Timestamp metav1.Time         `json:"timestamp"`
	Window    metav1.Duration     `json:"window"`
	Usage     corev1.ResourceList `json:"usage"`
}

func (obj *NodeMetrics) DeepCopyObject() runtime.Object {
	panic("implement me")
}

type NodeMetricsList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []NodeMetrics
}

func (obj *NodeMetricsList) DeepCopyObject() runtime.Object {
	panic("implement me")
}

type PodMetrics struct {
	metav1.TypeMeta
	metav1.ObjectMeta `json:"metadata"`
	Timestamp         metav1.Time     `json:"timestamp"`
	Window            metav1.Duration `json:"window"`

	Containers []metrics.ContainerMetrics `json:"containers"`
}

func (obj *PodMetrics) DeepCopyObject() runtime.Object {
	panic("implement me")
}

type PodMetricsList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []PodMetrics
}

func (obj *PodMetricsList) DeepCopyObject() runtime.Object {
	panic("implement me")
}
