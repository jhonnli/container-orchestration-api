package harbor

import "github.com/jhonnli/container-orchestration-api/model/harbor"

type ProjectInterface interface {
	List(env string) ([]harbor.Project, error)

	Get(env, projectId string) (*harbor.Project, error)
}
