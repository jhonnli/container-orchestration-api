package k8s

type PortParam struct {
	Port       int32 `json:"port" validate:"required"`
	TargetPort int32 `json:"targetPort" validate:"required"`
}

type ServiceSpecParam struct {
	Ports    []PortParam       `json:"ports" validate:"required,dive,gt=0"`
	Selector map[string]string `json:"selector" validate:"required,dive,gt=0"`
	Type     string            `json:"type"`
}

type ServiceParam struct {
	MetaData ObjectMetaParam  `json:"metaData" validate:"required,dive,required"`
	Spec     ServiceSpecParam `json:"spec" validate:"required,dive"`
}
