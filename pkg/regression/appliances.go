// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"net/http"

	openapiclient "cfm/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

var _ = DescribeRegression("Appliance Testing", func(sc *TestContext) {

	var (
		apiClient         *openapiclient.APIClient
		applianceId       string
		applianceCustomId string
	)

	BeforeEach(func() {

		By("setup cfm-service client for testing")
		logger := klog.FromContext(sc.Config.Ctx)
		logger.V(1).Info("config", "cfm ip", sc.Config.CFMService.IpAddress, "cfm port", sc.Config.CFMService.Port)
		apiClient = openapiclient.NewAPIClient(sc.Config.Config)

		Expect(apiClient).NotTo(BeNil())

		applianceCustomId = "ApplianceCustomId"
	})

	Describe("Appliances", func() {

		It("should have no appliances initially", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			appliances, httpRes, err := apiClient.DefaultAPI.AppliancesGet(sc.Config.Ctx).Execute()

			logger.V(3).Info("At Beginning", "appliances member count", appliances.GetMemberCount())

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(appliances).NotTo(BeNil())
			Expect(appliances.GetMembers()).To(BeNil())

		})
		It("should successfully add one appliance", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			request := apiClient.DefaultAPI.AppliancesPost(sc.Config.Ctx)

			applianceCredentials := openapiclient.Credentials{
				Username:  sc.Config.ApplianceEndpoints[0].Username,
				Password:  sc.Config.ApplianceEndpoints[0].Password,
				IpAddress: sc.Config.ApplianceEndpoints[0].IpAddress,
				Port:      sc.Config.ApplianceEndpoints[0].Port,
				Insecure:  &sc.Config.ApplianceEndpoints[0].Insecure,
				Protocol:  &sc.Config.ApplianceEndpoints[0].Protocol,
			}

			request = request.Credentials(applianceCredentials)

			appliance, httpRes, err := request.Execute()

			logger.V(3).Info("Add One Appliance", "appliance id", appliance.GetId())

			applianceId = appliance.GetId()

			Expect(err).To(BeNil())
			Expect(httpRes.StatusCode).To(Equal(http.StatusCreated))
			Expect(appliance).NotTo(BeNil())
		})

		It("should successfully show the list of all appliance ids", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			appliances, httpRes, err := apiClient.DefaultAPI.AppliancesGet(sc.Config.Ctx).Execute()

			logger.V(3).Info("Show All Appliance(s)", "MemberCount", appliances.GetMemberCount())

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(appliances).NotTo(BeNil())
			Expect(appliances.GetMemberCount()).NotTo(Equal(0))
		})

		It("should successfully show a particular appliance by id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			appliance, httpRes, err := apiClient.DefaultAPI.AppliancesGetById(sc.Config.Ctx, applianceId).Execute()

			logger.V(3).Info("Show One Appliance", "appliance id", appliance.GetId())

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(appliance).NotTo(BeNil())

		})

		It("should successfully delete an appliance by id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			appliance, httpRes, err := apiClient.DefaultAPI.AppliancesDeleteById(sc.Config.Ctx, applianceId).Execute()

			logger.V(3).Info("Delete One Appliance", "appliance id", appliance.GetId())

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(appliance).NotTo(BeNil())

		})

		It("should have no appliances after deleting the appliance", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			appliancesAfter, httpRes, err := apiClient.DefaultAPI.AppliancesGet(sc.Config.Ctx).Execute()

			logger.V(3).Info("At Last", "appliance member count", appliancesAfter.GetMemberCount())

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(appliancesAfter.GetMemberCount()).Should(BeZero())

		})

		It("should have no appliances", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			appliances, httpRes, err := apiClient.DefaultAPI.AppliancesGet(sc.Config.Ctx).Execute()

			logger.V(3).Info("At Beginning", "appliances member count", appliances.GetMemberCount())

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(appliances).NotTo(BeNil())
			Expect(appliances.GetMembers()).To(BeNil())

		})

		It("should successfully add one appliance with custom id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			request := apiClient.DefaultAPI.AppliancesPost(sc.Config.Ctx)

			applianceCredentials := openapiclient.Credentials{
				Username:  sc.Config.ApplianceEndpoints[0].Username,
				Password:  sc.Config.ApplianceEndpoints[0].Password,
				IpAddress: sc.Config.ApplianceEndpoints[0].IpAddress,
				Port:      sc.Config.ApplianceEndpoints[0].Port,
				Insecure:  &sc.Config.ApplianceEndpoints[0].Insecure,
				Protocol:  &sc.Config.ApplianceEndpoints[0].Protocol,
				CustomId:  &applianceCustomId,
			}

			request = request.Credentials(applianceCredentials)

			appliance, httpRes, err := request.Execute()

			logger.V(3).Info("Add One Appliance", "appliance id", appliance.GetId())

			applianceId = appliance.GetId()

			Expect(err).To(BeNil())
			Expect(httpRes.StatusCode).To(Equal(http.StatusCreated))
			Expect(appliance).NotTo(BeNil())
			Expect(applianceId).To(Equal(applianceCustomId))
		})

		It("should successfully show a particular appliance by custom id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			appliance, httpRes, err := apiClient.DefaultAPI.AppliancesGetById(sc.Config.Ctx, applianceId).Execute()

			applianceId = appliance.GetId()

			logger.V(3).Info("Show One Appliance", "appliance id", applianceId)

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(appliance).NotTo(BeNil())
			Expect(applianceId).To(Equal(applianceCustomId))

		})

		It("should successfully delete an appliance by id with custom id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			appliance, httpRes, err := apiClient.DefaultAPI.AppliancesDeleteById(sc.Config.Ctx, applianceId).Execute()

			applianceId = appliance.GetId()

			logger.V(3).Info("Delete One Appliance", "appliance id", applianceId)

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(appliance).NotTo(BeNil())
			Expect(applianceId).To(Equal(applianceCustomId))

		})

		It("should have no appliances after deleting the appliance", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			appliancesAfter, httpRes, err := apiClient.DefaultAPI.AppliancesGet(sc.Config.Ctx).Execute()

			logger.V(3).Info("At Last", "appliance member count", appliancesAfter.GetMemberCount())

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(appliancesAfter.GetMemberCount()).Should(BeZero())

		})

	})

})
