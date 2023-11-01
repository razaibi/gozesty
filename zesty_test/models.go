package main

// Define a struct to hold the YAML configuration
type TestConfig struct {
	Endpoints []EndpointConfig `yaml:"endpoints"`
}

type EndpointConfig struct {
	Name               string              `yaml:"name"`
	Parallel           bool                `yaml:"parallel"`
	EndpointURL        string              `yaml:"endpoint_url"`
	RequestMethod      string              `yaml:"request_method"`
	RequestHeaders     map[string]string   `yaml:"request_headers"`
	RequestBody        string              `yaml:"request_body"`
	ExpectedConditions []ExpectedCondition `yaml:"expected_conditions"`
}

type ExpectedCondition struct {
	ExpectedStatus   int    `yaml:"expected_status"`
	ExpectedResponse string `yaml:"expected_response"`
}
