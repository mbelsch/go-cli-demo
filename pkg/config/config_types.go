package config

type Config struct {
	Azure struct {
		CloudName        string `yaml:"cloudName"`
		TenantID         string `yaml:"tenantId"`
		SubscriptionID   string `yaml:"subscriptionId"`
		ResourceGroup    string `yaml:"resourceGroup"`
		ServicePrincipal struct {
			AppID string `yaml:"appId"`
		} `yaml:"servicePrincipal"`
		CertPath string `yaml:"certPath"`
		Password string `yaml:"password"`
	}
}
