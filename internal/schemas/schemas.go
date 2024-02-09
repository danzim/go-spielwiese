package schemas

// Status repräsentiert den Status einer Anfrage
type Status struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Status     string `json:"status"`
	Details    struct {
		Name string `json:"name"`
		Kind string `json:"kind"`
	} `json:"details"`
}

// ClusterList repräsentiert eine Liste von Clustern
type ClusterList struct {
	Kind       string    `json:"kind"`
	APIVersion string    `json:"apiVersion"`
	Items      []Cluster `json:"items"`
}

// Cluster repräsentiert einen Cluster
type Cluster struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name string `json:"name"`
	} `json:"metadata"`
	// Hier könntest du weitere Felder hinzufügen, die in der OpenAPI-Spezifikation definiert sind
}

// ProjectList repräsentiert eine Liste von Projekten
type ProjectList struct {
	Kind       string    `json:"kind"`
	APIVersion string    `json:"apiVersion"`
	Items      []Project `json:"items"`
}

// Project repräsentiert ein Projekt
type Project struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Spec struct {
		CI            string `json:"ci"`
		DisplayName   string `json:"displayName"`
		Description   string `json:"description"`
		CIApplication string `json:"ciApplication"`
		IONumber      string `json:"ioNumber"`
		PLKID         string `json:"plkID"`
		EgressIP      bool   `json:"egressIP"`
		// Hier könntest du weitere Felder hinzufügen, die in der OpenAPI-Spezifikation definiert sind
	} `json:"spec"`
}

type APIconfig struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
	} `yaml:"redis"`

	GRPC struct {
		BackendServiceAddress string `yaml:"backend_service_address"`
	} `yaml:"grpc"`

	Logging struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"logging"`
}
