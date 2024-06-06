// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"net/http"
	"strconv"
	"strings"

	openapiclient "cfm/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

var _ = DescribeRegression("ServiceVersion Testing", func(sc *TestContext) {

	var (
		apiClient *openapiclient.APIClient
	)

	BeforeEach(func() {

		By("setup cfm-service client for testing")
		logger := klog.FromContext(sc.Config.Ctx)
		logger.V(1).Info("config", "cfm ip", sc.Config.CFMService.IpAddress, "cfm port", sc.Config.CFMService.Port)
		apiClient = openapiclient.NewAPIClient(sc.Config.Config)

		Expect(apiClient).NotTo(BeNil())
	})

	Describe("ServiceVersion", func() {

		It("should get the service version and validate it's format", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			serviceInfo, httpRes, err := apiClient.DefaultAPI.CfmV1Get(sc.Config.Ctx).Execute()

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(serviceInfo).NotTo(BeNil())

			version := serviceInfo.GetVersion()
			logger.V(3).Info("cfm-service info", "version", version)
			Expect(len(version)).NotTo(BeZero())

			tokens := strings.Split(version, ".")
			Expect(len(tokens)).To(Equal(3)) //expect X.X.X format

			zerosFound := 0
			for _, token := range tokens {
				ver, err := strconv.Atoi(token)
				Expect(err).NotTo(HaveOccurred())
				Expect(ver).To(BeNumerically(">=", 0))
				if ver == 0 {
					zerosFound++
				}
			}

			Expect(zerosFound).ToNot(Equal(3))

		})

	})

})
