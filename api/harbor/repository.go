package harbor

import "github.com/jhonnli/container-orchestration-api/model/harbor"

type RepositoryInterface interface {
	List(env string, projectId string) ([]harbor.Repository, error)
	ListTag(env, projectId, repoName string) ([]harbor.DetailedTag, error)
}
