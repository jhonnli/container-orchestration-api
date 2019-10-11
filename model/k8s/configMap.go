package k8s

type ConfigMapParam struct {
	Name string            `json:"name"` //configMap的名称
	Data map[string]string `json:"data"` //configMap的内容
}
