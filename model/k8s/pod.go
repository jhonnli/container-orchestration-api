package k8s

type PodHealth struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Phase     string `json:"phase"`
	Reason    string `json:"reason"`
}

type PodHealthResponse struct {
	Check  bool        `json:"check"`
	Detail []PodHealth `json:"detail"`
}

type PodLogOptionParam struct {
	Container    string `form:"container"`
	SinceSeconds int64  `form:"sinceSeconds"`
	TailLines    int64  `form:"tailLines"`
	Timestamps   bool   `form:"timestamps"`
	Previous     bool   `form:"previous"`
	LimitBytes   int64  `form:"limitBytes"`
}
