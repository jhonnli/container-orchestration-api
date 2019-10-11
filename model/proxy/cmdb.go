package proxy

type KubernetesAuthInfo struct {
	Master string
	Token  string
}

type HarborAuthInfo struct {
	Server   string
	Username string
	Password string
}
