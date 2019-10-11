package k8s

import (
	corev1 "k8s.io/api/core/v1"
)

type EventInterface interface {
	Get(env, namespace, resourceName string) ([]corev1.Event, error)
}
