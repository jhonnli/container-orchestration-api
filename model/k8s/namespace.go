package k8s

type NamespaceParam struct {
	Name string `json:"name" validate:"required,max=30"`
}
