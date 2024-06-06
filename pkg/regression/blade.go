// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"net/http"

	openapiclient "cfm/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

var _ = DescribeRegression("Blade Testing", func(sc *TestContext) {

	var (
		apiClient     *openapiclient.APIClient
		applianceId   string
		bladeCustomId string
	)

	Describe("Blades", Ordered, func() {

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

			bladeCustomId = "BladeCustomId"

		})

		It("should have no blades initially", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			blades, response, err := apiClient.DefaultAPI.BladesGet(sc.Config.Ctx, applianceId).Execute()

			logger.V(3).Info("At Beginning", "blades member count", blades.GetMemberCount())

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(blades).NotTo(BeNil())
			Expect(blades.GetMembers()).To(BeNil())

		})
		It("should successfully added one blade", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			requestOfBlade := apiClient.DefaultAPI.BladesPost(sc.Config.Ctx, applianceId)

			credentials := openapiclient.Credentials{
				Username:  sc.Config.ApplianceEndpoints[0].Username,
				Password:  sc.Config.ApplianceEndpoints[0].Password,
				IpAddress: sc.Config.ApplianceEndpoints[0].IpAddress,
				Port:      sc.Config.ApplianceEndpoints[0].Port,
				Insecure:  &sc.Config.ApplianceEndpoints[0].Insecure,
				Protocol:  &sc.Config.ApplianceEndpoints[0].Protocol,
				CustomId:  &bladeCustomId,
			}

			requestOfBlade = requestOfBlade.Credentials(credentials)

			blade, response, err := requestOfBlade.Execute()

			actualBladeId := blade.GetId()

			logger.V(3).Info("Add One Blade", "blade id", actualBladeId)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusCreated))
			Expect(blade).NotTo(BeNil())
			Expect(actualBladeId).To(Equal(bladeCustomId))
		})

		It("should successfully show the list of all blade ids", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			blades, response, err := apiClient.DefaultAPI.BladesGet(sc.Config.Ctx, applianceId).Execute()

			logger.V(3).Info("Show All blade(s)", "MemberCount", blades.GetMemberCount())

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(blades).NotTo(BeNil())
			Expect(blades.GetMemberCount()).NotTo(Equal(0))

			members := blades.GetMembers()
			firstMember := members[0]
			Expect(firstMember.GetUri()).To(ContainSubstring("/cfm/v1/appliances/" + applianceId + "/blades/"))
		})

		It("should successfully show a particular blade by id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			blade, response, err := apiClient.DefaultAPI.BladesGetById(sc.Config.Ctx, applianceId, bladeCustomId).Execute()

			actualBladeId := blade.GetId()

			logger.V(3).Info("Show One Blade", "blade id", actualBladeId)

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(blade).NotTo(BeNil())
			Expect(actualBladeId).To(Equal(bladeCustomId))

		})

		It("should successfully delete a blade by id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			blade, response, err := apiClient.DefaultAPI.BladesDeleteById(sc.Config.Ctx, applianceId, bladeCustomId).Execute()

			actualBladeId := blade.GetId()

			logger.V(3).Info("Delete One Blade", "blade id", actualBladeId)

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(blade).NotTo(BeNil())
			Expect(actualBladeId).To(Equal(bladeCustomId))

		})

		It("should have no blades after deleting the blade", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			bladesAfter, response, err := apiClient.DefaultAPI.BladesGet(sc.Config.Ctx, applianceId).Execute()

			logger.V(3).Info("At Last", "appliance member count", bladesAfter.GetMemberCount())

			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
			Expect(bladesAfter.GetMemberCount()).Should(BeZero())

		})

		AfterAll(func() {

			logger := klog.FromContext(sc.Config.Ctx)

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
