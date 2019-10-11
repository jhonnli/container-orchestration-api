package k8s

import (
	"github.com/jhonnli/container-orchestration-api/model/k8s"
	"k8s.io/api/autoscaling/v2beta1"
)

type HPAInterface interface {
	Apply(env string, param k8s.HPAParam) (*v2beta1.HorizontalPodAutoscaler, error)
	Create(env string, param k8s.HPAParam) (*v2beta1.HorizontalPodAutoscaler, error)
	List(env string, namespace string) (*v2beta1.HorizontalPodAutoscalerList, error)
	Get(env string, namespace, name string) (*v2beta1.HorizontalPodAutoscaler, error)
	Delete(env string, namespace, name string) error
	Update(env string, param k8s.HPAParam) (*v2beta1.HorizontalPodAutoscaler, error)
}
