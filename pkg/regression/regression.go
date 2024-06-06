// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"bytes"
	"context"
	"fmt"

	openapiclient "cfm/pkg/client"

	"github.com/go-logr/logr"
	"github.com/spf13/viper"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// TestEndpoint provides configuration credentials for a specific web service,
// such as cfm-service, cxl-host, or a memory appliance.
type TestEndpoint struct {
	Username  string // The username credentials used to communicate with another service
	Password  string // The password credentials used to communicate with another service
	IpAddress string // The IP Address in dot notation of the service
	Port      int32  // The Port of the service
	Insecure  bool   // Set to true if an insecure connection should be used
	Protocol  string // Examples of http vs https
}

// TestConfig provides the configuration for the regression tests. It must be
// constructed with NewTestConfig to initialize it with sane defaults. The
// user of the regression package can then override values before passing
// the instance to [Ginkgo]Test and/or (when using GinkgoTest) in a
// BeforeEach. For example, the BeforeEach could set up the CFM Service
// and then set the CFMServiceIPAddress field differently for each test.
type TestConfig struct {
	CFMService         TestEndpoint                 // The CFM Service to use for testing
	Ctx                context.Context              // The cfm-regression context
	ApplianceEndpoints []TestEndpoint               // The Memory Appliances to use for testing
	CxlHostEndpoints   []TestEndpoint               // The CXL Hosts to use for testing
	Config             *openapiclient.Configuration // The OpenAPI Client configuration data
}

// TestContext gets initialized by the regression package before each test
// runs. It holds the variables that each test can depend on.
type TestContext struct {
	Config *TestConfig
}

// NewTestConfig returns a config instance with all values set to
// their defaults.
func NewTestConfig(ctx context.Context, configFile []byte) (TestConfig, error) {
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(configFile))

	config := TestConfig{}
	viper.Unmarshal(&config)

	if len(config.ApplianceEndpoints) == 0 || len(config.CxlHostEndpoints) == 0 {
		return config, fmt.Errorf("Configuration file must contain at least one ApplianceEndpoints[%d] and one CxlHostEndpoints[%d]",
			len(config.ApplianceEndpoints), len(config.CxlHostEndpoints))
	}

	config.Ctx = ctx

	config.Config = &openapiclient.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "CFM Service",
		Debug:         false,
		Servers: openapiclient.ServerConfigurations{
			{
				URL: fmt.Sprintf("http://%s:%d", config.CFMService.IpAddress, config.CFMService.Port),
			},
		},
		OperationServers: map[string]openapiclient.ServerConfigurations{},
	}

	return config, nil
}

// NewContext sets up regression testing with a config supplied by the
// user of the regression package. Ownership of that config is shared
// between the regression package and the caller.
func NewTestContext(config *TestConfig) *TestContext {
	return &TestContext{
		Config: config,
	}
}

// Test will test the CFM Service at the specified address by
// setting up a Ginkgo suite and running it.
func Test(t GinkgoTestingT, config TestConfig, logger logr.Logger) {
	sc := GinkgoTest(&config)
	RegisterFailHandler(Fail)

	RunSpecs(t, "CFM Service Regression Test Suite")
	sc.Finalize()
}

// GinkoTest is another entry point for regression testing: instead of
// directly running tests like Test does, it merely registers the
// tests. This can be used to embed regression testing in a custom Ginkgo
// test suite.  The pointer to the configuration is merely stored by
// GinkgoTest for use when the tests run. Therefore its content can
// still be modified in a BeforeEach. The regression package itself treats
// it as read-only.
func GinkgoTest(config *TestConfig) *TestContext {
	sc := NewTestContext(config)
	registerTestsInGinkgo(sc)
	return sc
}

// Setup must be invoked before each test. It initialize per-test
// variables in the context.
func (sc *TestContext) Setup() {
	// TODO
}

// Teardown must be called after each test. It frees resources
// allocated by Setup.
func (sc *TestContext) Teardown() {
	// TODO
}

// Finalize frees any resources that might be still cached in the context.
// It should be called after running all tests.
func (sc *TestContext) Finalize() {
	// TODO
}
