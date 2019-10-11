package k8s

type IngressParam struct {
	MetaData ObjectMetaParam  `json:"metaData"  validate:"required,dive"`
	Spec     IngressSpecParam `json:"spec"  validate:"required,dive"`
}

type IngressSpecParam struct {
	TLS   []IngressTLSParam `json:"tls,omitempty"`
	Rules []RuleParam       `json:"rules" validate:"required,dive"`
}

type RuleParam struct {
	Host string    `json:"host" validate:"required"`
	Http HttpParam `json:"http" validate:"required,dive"`
}

type HttpParam struct {
	Paths []PathParam `json:"paths" validate:"required,dive"`
}

type PathParam struct {
	Backend BackendParam `json:"backend" validate:"required,dive"`
	Path    string       `json:"path" validate:"required"`
}

type BackendParam struct {
	ServiceName string `json:"serviceName" validate:"required"`
	ServicePort int32  `json:"servicePort" validate:"required"`
}

type IngressTLSParam struct {
	Hosts      []string `json:"hosts,omitempty"`
	SecretName string   `json:"secretName,omitempty"`
}
