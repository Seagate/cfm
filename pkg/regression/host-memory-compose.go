// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"math"
	"net/http"
	"strings"

	openapiclient "cfm/pkg/client"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

// WARNING:
// Use CAUTION when enabling these tests.
// The cxl-host requires a reboot between compose and free actions for the memory changes to be visible to the cxl-host
// As such, these tests will only run under specific conditions
// These are rudementry tests that are testing very basic functionality.
// Even so, these test have been disabled when commited to the repo (to prevent ci\cd breakage)

var _ = DescribeRegression("Host Compose\\Free Memory Testing", func(sc *TestContext) {

	const (
		SIZE_OF_ONE_RESOURCE_MIB = int32(8192)
	)

	var (
		apiClient      *openapiclient.APIClient
		applianceId    string                         // Set during setup
		bladeId        string                         // Set during setup
		bladeMemoryId  = "memorychunk0"               // Just a default value.  May have to set manually within individual tests that are run standalone
		bladePortId    = "port0"                      // Just a default value.  May have to set manually within individual tests that are run standalone
		composeSizeMib = SIZE_OF_ONE_RESOURCE_MIB * 2 // Just a default value.  May have to set manually within individual tests that are run standalone
		qos            = openapiclient.Qos(4)         // Just a default value.  May have to set manually within individual tests that are run standalone
		// memoryCountBefore int32
		// memoryCountAfter  int32

		hostId       string
		hostPortId   = "port19-00"
		hostMemoryId = "node2"
	)

	Describe("HostMemComposition", Ordered, func() {

		BeforeAll(func() {

			logger := klog.FromContext(sc.Config.Ctx)

			By("setup cfm-service client")

			logger.V(1).Info("config", "cfm ip", sc.Config.CFMService.IpAddress, "cfm port", sc.Config.CFMService.Port)
			apiClient = openapiclient.NewAPIClient(sc.Config.Config)

			Expect(apiClient).NotTo(BeNil())

			By("setup host")

			hosts, response, err := apiClient.DefaultAPI.HostsGet(sc.Config.Ctx).Execute()

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(hosts).NotTo(BeNil())
			Expect(hosts.GetMembers()).To(BeNil())
			logger.V(3).Info("verified no hosts detected", "hosts member count", hosts.GetMemberCount())

			requestHost := apiClient.DefaultAPI.HostsPost(sc.Config.Ctx)
			credentialsHost := openapiclient.Credentials{
				Username:  sc.Config.CxlHostEndpoints[0].Username,
				Password:  sc.Config.CxlHostEndpoints[0].Password,
				IpAddress: sc.Config.CxlHostEndpoints[0].IpAddress,
				Port:      sc.Config.CxlHostEndpoints[0].Port,
				Insecure:  &sc.Config.CxlHostEndpoints[0].Insecure,
				Protocol:  &sc.Config.CxlHostEndpoints[0].Protocol,
			}

			requestHost = requestHost.Credentials(credentialsHost)
			host, response, err := requestHost.Execute()
			hostId = host.GetId()

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusCreated))
			Expect(host).NotTo(BeNil())
			logger.V(3).Info("added one host", "hostId", hostId)

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

			By("finish setup")
		})

		Describe("HostMemCompose", Ordered, func() {

			It("should get all memory (and save initial memory count)", Label("get"), func() {

				logger := klog.FromContext(sc.Config.Ctx)
				logger.V(3).Info("get memory", "hostId", hostId)

				memory, httpResponse, err := apiClient.DefaultAPI.HostGetMemory(sc.Config.Ctx, hostId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())

				for _, member := range memory.GetMembers() {
					uri := member.GetUri()
					elements := strings.Split(uri, "/")
					length := len(elements)
					Expect(length).NotTo(BeZero())
					Expect(elements[length-1]).To(ContainSubstring("node"))
					logger.V(3).Info("get memory", "uri", uri)
				}

				logger.V(3).Info("get memory", "memory count before compose", memory.GetMemberCount())
			})

			It("should compose a memory chunk", Label("host-compose"), func() {

				logger := klog.FromContext(sc.Config.Ctx)

				requestHost := apiClient.DefaultAPI.HostsComposeMemory(sc.Config.Ctx, hostId)
				composeRequest := openapiclient.NewComposeMemoryRequest(composeSizeMib, qos)
				composeRequest.SetPort(hostPortId)
				requestHost = requestHost.ComposeMemoryRequest(*composeRequest)
				memory, httpResponse, err := requestHost.Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusCreated))
				Expect(memory).NotTo(BeNil())
				Expect(memory.Id).To(Equal(bladeMemoryId)) // Because host's compose\free return the memory region as defined by the blade
				Expect(memory.MemoryApplianceId).NotTo(BeNil())
				Expect(*memory.MemoryApplianceId).To(Equal(applianceId))
				Expect(memory.MemoryBladeId).NotTo(BeNil())
				Expect(*memory.MemoryBladeId).To(Equal(bladeId))
				Expect(memory.Id).NotTo(BeEmpty())
				Expect(memory.MemoryAppliancePort).NotTo(BeNil())
				Expect(*memory.MemoryAppliancePort).To(Equal(bladePortId))

				bladeMemoryId = memory.Id
				memorySize := memory.SizeMiB

				Expect(memory.MappedHostId).NotTo(BeNil())
				Expect(*memory.MappedHostId).To(Equal(hostId))
				Expect(memory.MappedHostPort).NotTo(BeNil())
				Expect(*memory.MappedHostPort).To(Equal(hostPortId))

				logger.V(3).Info("compose memory", "hostId", hostId, "bladeMemoryId", bladeMemoryId, "memorySize(MiB)", memorySize)

				adjustMemorySize := int32((math.Ceil((math.Ceil(float64(composeSizeMib/SIZE_OF_ONE_RESOURCE_MIB)))/4))*4) * SIZE_OF_ONE_RESOURCE_MIB

				Expect(memorySize).To(Equal(adjustMemorySize))
				Expect(composeSizeMib).Should(BeNumerically("<=", memorySize))
			})

		})

		Describe("HostMemFree", Ordered, func() {

			It("should get all memory", Label("get"), func() {

				logger := klog.FromContext(sc.Config.Ctx)

				logger.V(3).Info("get memory", "hostId", hostId)

				memory, httpResponse, err := apiClient.DefaultAPI.HostGetMemory(sc.Config.Ctx, hostId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())

				for _, member := range memory.GetMembers() {
					uri := member.GetUri()
					elements := strings.Split(uri, "/")
					length := len(elements)
					Expect(length).NotTo(BeZero())
					Expect(elements[length-1]).To(ContainSubstring("node"))
					logger.V(3).Info("get memory", "uri", uri)
				}

				logger.V(3).Info("get memory", "count", memory.GetMemberCount())
			})

			It("should get a specific memory by id", Label("get"), func() {

				logger := klog.FromContext(sc.Config.Ctx)

				logger.V(3).Info("get memory by id", "hostId", hostId, "hostMemoryId", hostMemoryId)

				memory, httpResponse, err := apiClient.DefaultAPI.HostsGetMemoryById(sc.Config.Ctx, hostId, hostMemoryId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())
				Expect(memory.MappedHostId).NotTo(BeNil())
				Expect(*memory.MappedHostId).To(Equal(hostId))
				Expect(memory.MappedHostPort).NotTo(BeNil())
				Expect(*memory.MappedHostPort).To(Equal(hostPortId))
				Expect(memory.Id).To(Equal(hostMemoryId))
			})

			It("should delete a specific memory by id", Label("host-free"), func() {

				logger := klog.FromContext(sc.Config.Ctx)

				// if Label("free").MatchesLabelFilter(GinkgoLabelFilter()) {
				// 	// POTENTIAL USER INTERACTION.
				// 	// If running this spec standalone, need to manually update the memoryId.
				// 	logger.V(3).Info("delete memory by id: test parameter override")
				// 	hostMemoryId = "node2"
				// }

				logger.V(3).Info("delete memory by id", "hostId", hostId, "hostMemoryId", hostMemoryId)

				memory, httpResponse, err := apiClient.DefaultAPI.HostsFreeMemoryById(sc.Config.Ctx, hostId, hostMemoryId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())
				Expect(memory.Id).To(Equal(bladeMemoryId)) // Because host's compose\free return the memory region as defined by the blade
			})

			It("should get all memory", Label("get"), func() {

				logger := klog.FromContext(sc.Config.Ctx)

				logger.V(3).Info("get memory", "applianceId", applianceId, "bladeId", bladeId)

				memory, httpResponse, err := apiClient.DefaultAPI.BladesGetMemory(sc.Config.Ctx, applianceId, bladeId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())

				logger.V(3).Info("get memory", "memory count after FreeMemory", memory.GetMemberCount())
			})
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

			By("teardown host")
			host, response, err := apiClient.DefaultAPI.HostsDeleteById(sc.Config.Ctx, hostId).Execute()

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(host).NotTo(BeNil())
			logger.V(3).Info("deleted one host", "hostId", hostId)

			hosts, response, err := apiClient.DefaultAPI.HostsGet(sc.Config.Ctx).Execute()

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(hosts.GetMemberCount()).Should(BeZero())
			logger.V(3).Info("verified no hosts detected", "hosts member count", hosts.GetMemberCount())

		})

	})

})
