package k8s

import (
	model "github.com/jhonnli/container-orchestration-api/model/k8s"
)

type ApplicationInterface interface {
	Apply(env string, param model.ApplicationParam) map[string]bool
}
