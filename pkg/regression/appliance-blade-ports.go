// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"net/http"

	openapiclient "cfm/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

var _ = DescribeRegression("Blade Ports Testing", func(sc *TestContext) {

	var (
		apiClient   *openapiclient.APIClient
		applianceId string
		bladeId     string
		portId      = "port0"
	)

	Describe("BladePorts", Ordered, func() {

		BeforeAll(func() {

			logger := klog.FromContext(sc.Config.Ctx)

			By("setup cfm-service client")

			logger.V(1).Info("config", "cfm ip", sc.Config.CFMService.IpAddress, "cfm port", sc.Config.CFMService.Port)
			apiClient = openapiclient.NewAPIClient(sc.Config.Config)

			Expect(apiClient).NotTo(BeNil())

			By("setup appliance")

			appliances, response, err := apiClient.DefaultAPI.AppliancesGet(sc.Config.Ctx).Execute()

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(appliances).NotTo(BeNil())
			Expect(appliances.GetMembers()).To(BeNil())
			logger.V(3).Info("verified no appliances detected", "appliances member count", appliances.GetMemberCount())

			requestAppliance := apiClient.DefaultAPI.AppliancesPost(sc.Config.Ctx)
			//Note: Credentials no longer used internally by appliance, but still requried by service client.
			applianceCredentials := openapiclient.Credentials{
				Username:  sc.Config.ApplianceEndpoints[0].Username,
				Password:  sc.Config.ApplianceEndpoints[0].Password,
				IpAddress: sc.Config.ApplianceEndpoints[0].IpAddress,
				Port:      sc.Config.ApplianceEndpoints[0].Port,
				Insecure:  &sc.Config.ApplianceEndpoints[0].Insecure,
				Protocol:  &sc.Config.ApplianceEndpoints[0].Protocol,
			}
			requestAppliance = requestAppliance.Credentials(applianceCredentials)
			appliance, response, err := requestAppliance.Execute()
			applianceId = appliance.GetId()

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusCreated))
			Expect(appliance).NotTo(BeNil())
			logger.V(3).Info("added one appliance", "applianceId", applianceId)

			By("setup blade")

			blades, response, err := apiClient.DefaultAPI.BladesGet(sc.Config.Ctx, applianceId).Execute()

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(blades).NotTo(BeNil())
			Expect(blades.GetMembers()).To(BeNil())
			logger.V(3).Info("verified no blades detected", "blades member count", blades.GetMemberCount())

			requestBlade := apiClient.DefaultAPI.BladesPost(sc.Config.Ctx, applianceId)
			credentials := openapiclient.Credentials{
				Username:  sc.Config.ApplianceEndpoints[0].Username,
				Password:  sc.Config.ApplianceEndpoints[0].Password,
				IpAddress: sc.Config.ApplianceEndpoints[0].IpAddress,
				Port:      sc.Config.ApplianceEndpoints[0].Port,
				Insecure:  &sc.Config.ApplianceEndpoints[0].Insecure,
				Protocol:  &sc.Config.ApplianceEndpoints[0].Protocol,
			}
			requestBlade = requestBlade.Credentials(credentials)
			blade, response, err := requestBlade.Execute()
			bladeId = blade.GetId()

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusCreated))
			Expect(blade).NotTo(BeNil())
			logger.V(3).Info("added one blade", "bladeId", bladeId)

		})

		It("should get all ports", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			logger.V(3).Info("get ports", "applianceId", applianceId)
			logger.V(3).Info("get ports", "bladeId", bladeId)

			ports, httpResponse, err := apiClient.DefaultAPI.BladesGetPorts(sc.Config.Ctx, applianceId, bladeId).Execute()

			Expect(err).To(BeNil())
			Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
			Expect(ports).NotTo(BeNil())
			Expect(ports.GetMemberCount()).Should(BeNumerically("==", 4))
			members := ports.GetMembers()
			firstMember := members[0]
			Expect(firstMember.GetUri()).To(ContainSubstring("/cfm/v1/appliances/" + applianceId + "/blades/" + bladeId + "/ports/" + portId))

		})

		It("should get a specific port by id", func() {

			logger := klog.FromContext(sc.Config.Ctx)
			logger.V(3).Info("get port by id", "applianceId", applianceId)
			logger.V(3).Info("get port by id", "bladeId", bladeId)
			logger.V(3).Info("get port by id", "portId", portId)

			port, httpResponse, err := apiClient.DefaultAPI.BladesGetPortById(sc.Config.Ctx, applianceId, bladeId, portId).Execute()

			Expect(err).To(BeNil())
			Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
			Expect(port).NotTo(BeNil())

			Expect(port.GetId()).To(BeEquivalentTo(portId))
			Expect(port.GetGCxlId()).ShouldNot(HaveLen(0))
			Expect(port.GetPortProtocol()).To(BeEquivalentTo("CXL"))
			Expect(port.GetStatusHealth()).To(BeEquivalentTo("Ok"))

		})

		AfterAll(func() {
			logger := klog.FromContext(sc.Config.Ctx)

			By("teardown blade")
			blade, response, err := apiClient.DefaultAPI.BladesDeleteById(sc.Config.Ctx, applianceId, bladeId).Execute()

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(blade).NotTo(BeNil())
			logger.V(3).Info("deleted one blade", "bladeId", bladeId)

			blades, response, err := apiClient.DefaultAPI.BladesGet(sc.Config.Ctx, applianceId).Execute()

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(blades.GetMemberCount()).Should(BeZero())
			logger.V(3).Info("verified no blades detected", "blades member count", blades.GetMemberCount())

			By("teardown appliance")
			appliance, response, err := apiClient.DefaultAPI.AppliancesDeleteById(sc.Config.Ctx, applianceId).Execute()

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(appliance).NotTo(BeNil())
			logger.V(3).Info("deleted one appliance", "applianceId", applianceId)

			appliances, response, err := apiClient.DefaultAPI.AppliancesGet(sc.Config.Ctx).Execute()

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(appliances.GetMemberCount()).Should(BeZero())
			logger.V(3).Info("verified no appliances detected", "appliances member count", appliances.GetMemberCount())

		})

	})

})
