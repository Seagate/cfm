// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"math"
	"net/http"

	openapiclient "cfm/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

// TODO: Need to improve cause the uris changes
var _ = DescribeRegression("Appliance Blade Memory Testing", func(sc *TestContext) {

	const (
		SIZE_OF_ONE_RESOURCE_MIB = int32(8192)
	)

	var (
		apiClient         *openapiclient.APIClient
		applianceId       string                         // Set during setup
		bladeId           string                         // Set during setup
		memoryId          string                         // Only set after Compose...() API calls. May have to set manually within individual tests that are run standalone
		portId            = "port0"                      // Just a default value.  May have to set manually within individual tests that are run standalone
		composeSizeMib    = SIZE_OF_ONE_RESOURCE_MIB * 2 // Just a default value.  May have to set manually within individual tests that are run standalone
		qos               = openapiclient.Qos(4)         // Just a default value.  May have to set manually within individual tests that are run standalone
		memoryCountBefore int32
		memoryCountAfter  int32
	)

	Describe("BladeMemory", Ordered, func() {

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

			By("finish setup")
		})

		Describe("BladeMemoryProvision", Ordered, func() {

			It("should get all memory (and save initial memory count)", func() {
				logger := klog.FromContext(sc.Config.Ctx)
				logger.V(3).Info("get memory", "applianceId", applianceId, "bladeId", bladeId)

				memory, httpResponse, err := apiClient.DefaultAPI.BladesGetMemory(sc.Config.Ctx, applianceId, bladeId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())

				memoryCountBefore = memory.GetMemberCount()

				logger.V(3).Info("get memory", "memory count before provision", memoryCountBefore)
			})

			It("should provision a memory chunk", Label("allocate", "size"), func() {
				logger := klog.FromContext(sc.Config.Ctx)

				bladeRequest := apiClient.DefaultAPI.BladesComposeMemory(sc.Config.Ctx, applianceId, bladeId)
				composeRequest := openapiclient.NewComposeMemoryRequest(composeSizeMib, qos)
				bladeRequest = bladeRequest.ComposeMemoryRequest(*composeRequest)
				memory, httpResponse, err := bladeRequest.Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusCreated))
				Expect(memory).NotTo(BeNil())
				Expect(memory.MemoryApplianceId).NotTo(BeNil())
				Expect(*memory.MemoryApplianceId).To(Equal(applianceId))
				Expect(memory.MemoryBladeId).NotTo(BeNil())
				Expect(*memory.MemoryBladeId).To(Equal(bladeId))
				Expect(memory.Id).NotTo(BeEmpty())
				Expect(memory.MemoryAppliancePort).To(BeNil())

				memoryId = memory.Id
				memorySize := memory.SizeMiB

				logger.V(3).Info("compose memory by size", "applianceId", applianceId, "bladeId", bladeId, "memoryId", memoryId, "memorySize(MiB)", memorySize)

				adjustMemorySize := int32((math.Ceil((math.Ceil(float64(composeSizeMib/SIZE_OF_ONE_RESOURCE_MIB)))/4))*4) * SIZE_OF_ONE_RESOURCE_MIB

				Expect(memorySize).To(Equal(adjustMemorySize))
				Expect(composeSizeMib).Should(BeNumerically("<=", memorySize))
			})

			It("should provision by resource a memory chunk", Label("allocate", "resource"), func() {
				logger := klog.FromContext(sc.Config.Ctx)

				// resourceIds := []string{"resourceblock1", "resourceblock5", "resourceblock9", "resourceblock13"}
				resourceIds := []string{"resourceblock1", "resourceblock5", "resourceblock9", "resourceblock13"}

				bladeRequest := apiClient.DefaultAPI.BladesComposeMemoryByResource(sc.Config.Ctx, applianceId, bladeId)
				composeRequest := openapiclient.NewComposeMemoryByResourceRequest(resourceIds)
				bladeRequest = bladeRequest.ComposeMemoryByResourceRequest(*composeRequest)
				memory, httpResponse, err := bladeRequest.Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusCreated))
				Expect(memory).NotTo(BeNil())
				Expect(memory.MemoryApplianceId).NotTo(BeNil())
				Expect(*memory.MemoryApplianceId).To(Equal(applianceId))
				Expect(memory.MemoryBladeId).NotTo(BeNil())
				Expect(*memory.MemoryBladeId).To(Equal(bladeId))
				Expect(memory.Id).NotTo(BeEmpty())
				Expect(memory.MemoryAppliancePort).To(BeNil())

				memoryId = memory.Id
				memorySize := memory.SizeMiB

				logger.V(3).Info("compose memory by resource", "applianceId", applianceId, "bladeId", bladeId, "memoryId", memoryId, "memorySize(MiB)", memorySize)

				adjustMemorySize := int32(len(resourceIds)) * SIZE_OF_ONE_RESOURCE_MIB

				Expect(memorySize).To(Equal(adjustMemorySize))
				Expect(composeSizeMib).Should(BeNumerically("<=", memorySize))
			})

			It("should get all memory (and verify memory count increment)", func() {

				logger := klog.FromContext(sc.Config.Ctx)

				logger.V(3).Info("get memory", "applianceId", applianceId, "bladeId", bladeId)

				memory, httpResponse, err := apiClient.DefaultAPI.BladesGetMemory(sc.Config.Ctx, applianceId, bladeId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())

				memoryCountAfter = memory.GetMemberCount()
				Expect(memoryCountAfter).To(Equal(memoryCountBefore + 1))

				logger.V(3).Info("get memory", "memory count after provision", memoryCountAfter)
			})

			It("should get a specific memory by id", func() {

				logger := klog.FromContext(sc.Config.Ctx)

				logger.V(3).Info("get memory by id", "applianceId", applianceId, "bladeId", bladeId, "memoryId", memoryId)

				memory, httpResponse, err := apiClient.DefaultAPI.BladesGetMemoryById(sc.Config.Ctx, applianceId, bladeId, memoryId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())
				Expect(memory.MemoryApplianceId).NotTo(BeNil())
				Expect(*memory.MemoryApplianceId).To(Equal(applianceId))
				Expect(memory.MemoryBladeId).NotTo(BeNil())
				Expect(*memory.MemoryBladeId).To(Equal(bladeId))
				Expect(memory.Id).To(Equal(memoryId))
				Expect(memory.MemoryAppliancePort).To(BeNil())
			})

			It("should fail to assign a bad portId to a specific memory by id", func() {

				logger := klog.FromContext(sc.Config.Ctx)

				invalidPortId := "invalidPort"

				logger.V(3).Info("assign memory", "applianceId", applianceId, "bladeId", bladeId, "memoryId", memoryId, "portId", invalidPortId)

				bladeRequest := apiClient.DefaultAPI.BladesAssignMemoryById(sc.Config.Ctx, applianceId, bladeId, memoryId)
				assignRequest := openapiclient.NewAssignMemoryRequest(invalidPortId, "assign")
				bladeRequest = bladeRequest.AssignMemoryRequest(*assignRequest)
				memory, httpResponse, err := bladeRequest.Execute()

				Expect(err).NotTo(BeNil())
				Expect(httpResponse).NotTo(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusConflict))
				Expect(memory).To(BeNil())
			})

			It("should assign a port to a specific memory by id", Label("assign"), func() {

				logger := klog.FromContext(sc.Config.Ctx)

				labels := GinkgoLabelFilter()
				notAssign := Label("!assign").MatchesLabelFilter(labels)
				if !notAssign { // Yes, this is weird logic. When using a "assign" Label, "!notAssign" does not behave the same as "Label("assign").MatchesLabelFilter(labels)"
					// POTENTIAL USER INTERACTION.
					// ONLY want to enter here if running this spec standalone.
					// If so, may need to manually update the memoryId and\or portId here.
					logger.V(3).Info("assign memory: test parameter override")
					memoryId = "memorychunk0"
					portId = "port0"
				}

				logger.V(3).Info("assign memory", "applianceId", applianceId, "bladeId", bladeId, "memoryId", memoryId, "portId", portId)

				bladeRequest := apiClient.DefaultAPI.BladesAssignMemoryById(sc.Config.Ctx, applianceId, bladeId, memoryId)
				assignRequest := openapiclient.NewAssignMemoryRequest(portId, "assign")
				bladeRequest = bladeRequest.AssignMemoryRequest(*assignRequest)
				memory, httpResponse, err := bladeRequest.Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())
				Expect(memory.MemoryApplianceId).NotTo(BeNil())
				Expect(*memory.MemoryApplianceId).To(Equal(applianceId))
				Expect(memory.MemoryBladeId).NotTo(BeNil())
				Expect(*memory.MemoryBladeId).To(Equal(bladeId))
				Expect(memory.Id).To(Equal(memoryId))
				Expect(memory.MemoryAppliancePort).NotTo(BeNil())
				Expect(*memory.MemoryAppliancePort).To(Equal(portId))
			})

			It("should unassign a port for a specific memory by id", Label("unassign"), func() {

				logger := klog.FromContext(sc.Config.Ctx)

				labels := GinkgoLabelFilter()
				notUnassign := Label("!unassign").MatchesLabelFilter(labels)
				if !notUnassign { // Yes, this is weird logic. When using a "unassign" Label, "!notUnassign" does not behave the same as "Label("unassign").MatchesLabelFilter(labels)"
					// POTENTIAL USER INTERACTION.
					// ONLY want to enter here if running this spec standalone.
					// If so, may need to manually update the memoryId and\or portId here.
					logger.V(3).Info("unassign memory: test parameter override")
					memoryId = "memorychunk0"
					portId = "port0"
				}

				logger.V(3).Info("unassign memory", "applianceId", applianceId, "bladeId", bladeId, "memoryId", memoryId, "portId", portId)

				bladeRequest := apiClient.DefaultAPI.BladesAssignMemoryById(sc.Config.Ctx, applianceId, bladeId, memoryId)
				assignRequest := openapiclient.NewAssignMemoryRequest(portId, "unassign")
				bladeRequest = bladeRequest.AssignMemoryRequest(*assignRequest)
				memory, httpResponse, err := bladeRequest.Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())
				Expect(memory.MemoryApplianceId).NotTo(BeNil())
				Expect(*memory.MemoryApplianceId).To(Equal(applianceId))
				Expect(memory.MemoryBladeId).NotTo(BeNil())
				Expect(*memory.MemoryBladeId).To(Equal(bladeId))
				Expect(memory.Id).To(Equal(memoryId))
			})

			It("should delete a specific memory by id", Label("free"), func() {

				logger := klog.FromContext(sc.Config.Ctx)

				labels := GinkgoLabelFilter()
				notFree := Label("!free").MatchesLabelFilter(labels)
				if !notFree { // Yes, this is weird logic. When using a "free" Label, "!notFree" does not behave the same as "Label("free").MatchesLabelFilter(labels)"
					// POTENTIAL USER INTERACTION.
					// ONLY want to enter here if running this spec standalone.
					// If so, may need to manually update the memoryId here.
					logger.V(3).Info("delete memory by id: test parameter override")
					memoryId = "memorychunk0"
				}

				logger.V(3).Info("delete memory by id", "applianceId", applianceId, "bladeId", bladeId, "memoryId", memoryId)

				memory, httpResponse, err := apiClient.DefaultAPI.BladesFreeMemoryById(sc.Config.Ctx, applianceId, bladeId, memoryId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())
			})

			It("should get all memory (and verify memory count decrement)", func() {

				logger := klog.FromContext(sc.Config.Ctx)

				logger.V(3).Info("get memory", "applianceId", applianceId, "bladeId", bladeId)

				memory, httpResponse, err := apiClient.DefaultAPI.BladesGetMemory(sc.Config.Ctx, applianceId, bladeId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())
				// After the provisioned memory is deleted, the memory count returns to the value before provisioning.
				Expect(memory.GetMemberCount()).To(Equal(memoryCountBefore))

				logger.V(3).Info("get memory", "memory count after provision free", memoryCountBefore)
			})

		})

		Describe("BladeMemoryCompose", Ordered, func() {

			It("should get all memory (and save initial memory count)", func() {

				logger := klog.FromContext(sc.Config.Ctx)
				logger.V(3).Info("get memory", "applianceId", applianceId, "bladeId", bladeId)

				memory, httpResponse, err := apiClient.DefaultAPI.BladesGetMemory(sc.Config.Ctx, applianceId, bladeId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())

				memoryCountBefore = memory.GetMemberCount()

				logger.V(3).Info("get memory", "memory count before compose", memoryCountBefore)
			})

			It("should compose a memory chunk", Label("allocate", "size"), func() {

				logger := klog.FromContext(sc.Config.Ctx)

				bladeRequest := apiClient.DefaultAPI.BladesComposeMemory(sc.Config.Ctx, applianceId, bladeId)
				composeRequest := openapiclient.NewComposeMemoryRequest(composeSizeMib, qos)
				composeRequest.SetPort(portId)
				bladeRequest = bladeRequest.ComposeMemoryRequest(*composeRequest)
				memory, httpResponse, err := bladeRequest.Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusCreated))
				Expect(memory).NotTo(BeNil())
				Expect(memory.MemoryApplianceId).NotTo(BeNil())
				Expect(*memory.MemoryApplianceId).To(Equal(applianceId))
				Expect(memory.MemoryBladeId).NotTo(BeNil())
				Expect(*memory.MemoryBladeId).To(Equal(bladeId))
				Expect(memory.Id).NotTo(BeEmpty())
				Expect(memory.MemoryAppliancePort).NotTo(BeNil())
				Expect(*memory.MemoryAppliancePort).To(Equal(portId))

				memoryId = memory.Id
				memorySize := memory.SizeMiB

				logger.V(3).Info("compose memory", "applianceId", applianceId, "bladeId", bladeId, "memoryId", memoryId, "memorySize(MiB)", memorySize)

				adjustMemorySize := int32((math.Ceil((math.Ceil(float64(composeSizeMib/SIZE_OF_ONE_RESOURCE_MIB)))/4))*4) * SIZE_OF_ONE_RESOURCE_MIB

				Expect(memorySize).To(Equal(adjustMemorySize))
				Expect(composeSizeMib).Should(BeNumerically("<=", memorySize))
			})

			It("should provision by resource a memory chunk", Label("allocate", "resource"), func() {
				logger := klog.FromContext(sc.Config.Ctx)

				// resourceIds := []string{"resourceblock1", "resourceblock5", "resourceblock9", "resourceblock13"}
				resourceIds := []string{"resourceblock2", "resourceblock3", "resourceblock4", "resourceblock5"}

				bladeRequest := apiClient.DefaultAPI.BladesComposeMemoryByResource(sc.Config.Ctx, applianceId, bladeId)
				composeRequest := openapiclient.NewComposeMemoryByResourceRequest(resourceIds)
				composeRequest.SetPort(portId)
				bladeRequest = bladeRequest.ComposeMemoryByResourceRequest(*composeRequest)
				memory, httpResponse, err := bladeRequest.Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusCreated))
				Expect(memory).NotTo(BeNil())
				Expect(memory.MemoryApplianceId).NotTo(BeNil())
				Expect(*memory.MemoryApplianceId).To(Equal(applianceId))
				Expect(memory.MemoryBladeId).NotTo(BeNil())
				Expect(*memory.MemoryBladeId).To(Equal(bladeId))
				Expect(memory.Id).NotTo(BeEmpty())
				Expect(memory.MemoryAppliancePort).NotTo(BeNil())
				Expect(*memory.MemoryAppliancePort).To(Equal(portId))

				memoryId = memory.Id
				memorySize := memory.SizeMiB

				logger.V(3).Info("compose memory by resource", "applianceId", applianceId, "bladeId", bladeId, "memoryId", memoryId, "memorySize(MiB)", memorySize)

				adjustMemorySize := int32(len(resourceIds)) * SIZE_OF_ONE_RESOURCE_MIB

				Expect(memorySize).To(Equal(adjustMemorySize))
				Expect(composeSizeMib).Should(BeNumerically("<=", memorySize))
			})

			It("should get all memory (and verify memory count increment)", func() {

				logger := klog.FromContext(sc.Config.Ctx)

				logger.V(3).Info("get memory", "applianceId", applianceId, "bladeId", bladeId)

				memory, httpResponse, err := apiClient.DefaultAPI.BladesGetMemory(sc.Config.Ctx, applianceId, bladeId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())

				memoryCountAfter = memory.GetMemberCount()
				Expect(memoryCountAfter).To(Equal(memoryCountBefore + 1))

				logger.V(3).Info("get memory", "memory count after Compose Memory", memoryCountAfter)
			})

			It("should get a specific memory by id", func() {

				logger := klog.FromContext(sc.Config.Ctx)

				logger.V(3).Info("get memory by id", "applianceId", applianceId, "bladeId", bladeId, "memoryId", memoryId)

				memory, httpResponse, err := apiClient.DefaultAPI.BladesGetMemoryById(sc.Config.Ctx, applianceId, bladeId, memoryId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())
				Expect(memory.MemoryApplianceId).NotTo(BeNil())
				Expect(*memory.MemoryApplianceId).To(Equal(applianceId))
				Expect(memory.MemoryBladeId).NotTo(BeNil())
				Expect(*memory.MemoryBladeId).To(Equal(bladeId))
				Expect(memory.Id).To(Equal(memoryId))
				Expect(memory.MemoryAppliancePort).NotTo(BeNil())
				Expect(*memory.MemoryAppliancePort).To(Equal(portId))
			})

			It("should delete a specific memory by id", Label("free"), func() {

				logger := klog.FromContext(sc.Config.Ctx)

				if Label("free").MatchesLabelFilter(GinkgoLabelFilter()) {
					// POTENTIAL USER INTERACTION.
					// If running this spec standalone, need to manually update the memoryId.
					logger.V(3).Info("delete memory by id: test parameter override")
					memoryId = "memorychunk0"
				}

				logger.V(3).Info("delete memory by id", "applianceId", applianceId, "bladeId", bladeId, "memoryId", memoryId)

				memory, httpResponse, err := apiClient.DefaultAPI.BladesFreeMemoryById(sc.Config.Ctx, applianceId, bladeId, memoryId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())

			})

			It("should get all memory (and verify memory count decrement)", func() {

				logger := klog.FromContext(sc.Config.Ctx)

				logger.V(3).Info("get memory", "applianceId", applianceId, "bladeId", bladeId)

				memory, httpResponse, err := apiClient.DefaultAPI.BladesGetMemory(sc.Config.Ctx, applianceId, bladeId).Execute()

				Expect(err).To(BeNil())
				Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
				Expect(memory).NotTo(BeNil())
				// After the composed memory is deleted, the memory count returns to the value before composing.
				Expect(memory.GetMemberCount()).To(Equal(memoryCountBefore))

				logger.V(3).Info("get memory", "memory count after FreeMemory", memoryCountBefore)
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

		})

	})

})
