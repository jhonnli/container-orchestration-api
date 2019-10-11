package k8s

type HPAParam struct {
	MetaData ObjectMetaParam `json:"metaData" validate:"required,dive"`
	Spec     HPASpecParam    `json:"spec" validate:"required,dive"`
}

type HPASpecParam struct {
	MaxReplicas        int32           `json:"maxReplicas" validate:"required"`
	MinReplicas        int32           `json:"minReplicas" validate:"required"`
	Metrics            HPAMetricsParam `json:"metrics" validate:"required,dive"`
	ScaleTargetRefName string          `json:"scaleTargetRefName"`
}

type HPAMetricsParam struct {
	Cpu    int32 `json:"cpu" validate:"required,numeric"`
	Memory int32 `json:"memory" validate:"required",numeric`
}
