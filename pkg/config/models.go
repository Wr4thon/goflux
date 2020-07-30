package config

// Config is used to configure goflux
type Config struct {
	Deployment Deployment `json:"deployment,omitempty" yaml:"deployment,omitempty"`
	HPA        HPA        `json:"hpa,omitempty" yaml:"hpa,omitempty"`
	Secrets    Secrets    `json:"secrets,omitempty" yaml:"secrets,omitempty"`
}

// Secrets is used to seal and find secrets
type Secrets struct {
	SecretFolderName string `json:"secretFolderName,omitempty" yaml:"secretFolderName,omitempty"`
	DevCertURL       string `json:"devCertURL,omitempty" yaml:"devCertURL,omitempty"`
	TestCertURL      string `json:"testCertURL,omitempty" yaml:"testCertURL,omitempty"`
	ProdCertURL      string `json:"prodCertURL,omitempty" yaml:"prodCertURL,omitempty"`
}

// HPA is used to configure the creation of hpa files
type HPA struct {
	MinReplicas int `json:"minReplicas,omitempty" yaml:"minReplicas,omitempty"`
	MaxReplicas int `json:"maxReplicas,omitempty" yaml:"maxReplicas,omitempty"`
}

// Deployment is used to configure the creation of deployment files
type Deployment struct {
	ImagePullSecret string      `json:"imagePullSecret,omitempty" yaml:"imagePullSecret,omitempty"`
	Annotations     Annotations `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Ressources      Ressources  `json:"ressources,omitempty" yaml:"ressources,omitempty"`
}

// Ressources are used as ressources
type Ressources struct {
	Limits   Ressource `json:"limits,omitempty" yaml:"limits,omitempty"`
	Requests Ressource `json:"requests,omitempty" yaml:"requests,omitempty"`
}

// Ressource is used inside ressources
type Ressource struct {
	CPU    string `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`
}

// Annotations are used as annotations
type Annotations map[string]string
