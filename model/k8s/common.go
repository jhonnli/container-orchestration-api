package k8s

type ObjectMeta struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type ObjectMetaParam struct {
	Name        string            `json:"name" validate:"required"`
	Namespace   string            `json:"namespace" validate:"required"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type ListOptions struct {
	LabelSelector string `json:"labelSelector,omitempty"`

	FieldSelector string `json:"fieldSelector,omitempty"`

	IncludeUninitialized bool `json:"includeUninitialized"`

	Watch bool `json:"watch,omitempty"`

	ResourceVersion string `json:"resourceVersion,omitempty"`

	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`
	Limit          int64  `json:"limit,omitempty"`
	Continue       string `json:"continue,omitempty"`
}
