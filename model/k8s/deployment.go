package k8s

type ContainerParam struct {
	Name            string                  `json:"name"`
	Image           string                  `json:"image"`
	ServicePort     int32                   `json:"servicePort"`
	HealthCheckPath string                  `json:"healthCheckPath"`
	Resources       ContainerResourceParams `json:"resources"`
	VolumeMounts    []VolumeMountParam      `json:"volumeMounts"`
}

type DeploymentParam struct {
	MetaData ObjectMetaParam     `json:"metaData" validate:"required,dive"`
	Spec     DeploymentSpecParam `json:"spec"`
}

type DeploymentSpecParam struct {
	Replicas int32                   `json:"replicas"`
	Template DeploymentTemplateParam `json:"template"`
}

type DeploymentTemplateParam struct {
	Spec DeploymentTemplateSpecParam `json:"spec"`
}

type DeploymentTemplateSpecParam struct {
	Containers   []ContainerParam  `json:"containers"`
	NodeSelector map[string]string `json:"nodeSelector"`
}

type ContainerResourceParam struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}

type ContainerResourceParams struct {
	Limits   ContainerResourceParam `json:"limits"`
	Requests ContainerResourceParam `json:"requests"`
}

type VolumeMountParam struct {
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
	Name      string `json:"name"`
}
