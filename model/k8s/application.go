package k8s

type ApplicationParam struct {
	Deployment              DeploymentParam `json:"deployment" validate:"dive"`
	Service                 ServiceParam    `json:"service" validate:"dive"`
	Ingress                 IngressParam    `json:"ingress" validate:"dive"`
	HorizontalPodAutoscaler HPAParam        `json:"horizontalPodAutoscaler" validate:"dive"`
}
