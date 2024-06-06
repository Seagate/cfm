// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"net/http"

	openapiclient "cfm/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

var _ = DescribeRegression("Host Testing", func(sc *TestContext) {

	var (
		apiClient    *openapiclient.APIClient
		hostCustomId string
	)

	BeforeEach(func() {

		By("setup client for testing")
		apiClient = openapiclient.NewAPIClient(sc.Config.Config)

		hostCustomId = "HostCustomId"

	})

	Describe("Hosts", func() {

		It("should have no host initially", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			hosts, httpRes, err := apiClient.DefaultAPI.HostsGet(sc.Config.Ctx).Execute()

			logger.V(3).Info("At Beginning", "host member count", hosts.GetMemberCount())

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(hosts).NotTo(BeNil())
			Expect(hosts.GetMembers()).To(BeNil())

		})

		It("should successfully added one host", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			request := apiClient.DefaultAPI.HostsPost(sc.Config.Ctx)

			hostCredentials := openapiclient.Credentials{
				Username:  sc.Config.CxlHostEndpoints[0].Username,
				Password:  sc.Config.CxlHostEndpoints[0].Password,
				IpAddress: sc.Config.CxlHostEndpoints[0].IpAddress,
				Port:      sc.Config.CxlHostEndpoints[0].Port,
				Insecure:  &sc.Config.CxlHostEndpoints[0].Insecure,
				Protocol:  &sc.Config.CxlHostEndpoints[0].Protocol,
				CustomId:  &hostCustomId,
			}

			request = request.Credentials(hostCredentials)

			host, httpRes, err := request.Execute()

			actualHostId := host.GetId()

			logger.V(3).Info("Add One Host", "host id", actualHostId)

			Expect(err).To(BeNil())
			Expect(httpRes.StatusCode).To(Equal(http.StatusCreated))
			Expect(host).NotTo(BeNil())
			Expect(actualHostId).To(Equal(hostCustomId))

		})

		It("should successfully show the list of all host ids", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			hosts, httpRes, err := apiClient.DefaultAPI.HostsGet(sc.Config.Ctx).Execute()

			logger.V(3).Info("Show All Host(s)", "MemberCount", hosts.GetMemberCount())

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(hosts).NotTo(BeNil())
			Expect(hosts.GetMemberCount()).NotTo(Equal(0))
		})

		It("should successfully show a particular host by id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			host, httpRes, err := apiClient.DefaultAPI.HostsGetById(sc.Config.Ctx, hostCustomId).Execute()

			actualHostId := host.GetId()

			logger.V(3).Info("Show One Host", "host id", actualHostId)

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(host).NotTo(BeNil())
			Expect(actualHostId).To(Equal(hostCustomId))

		})

		It("should successfully delete a host by id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			host, httpRes, err := apiClient.DefaultAPI.HostsDeleteById(sc.Config.Ctx, hostCustomId).Execute()

			actualHostId := host.GetId()

			logger.V(3).Info("Deleted One Host", "host id", actualHostId)

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(host).NotTo(BeNil())
			Expect(actualHostId).To(Equal(hostCustomId))

		})

		It("should have no host after deleting the host", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			hostsAfter, httpRes, err := apiClient.DefaultAPI.HostsGet(sc.Config.Ctx).Execute()

			logger.V(3).Info("At Last", "host member count", hostsAfter.GetMemberCount())

			Expect(err).NotTo(HaveOccurred())
			Expect(httpRes.StatusCode).To(Equal(http.StatusOK))
			Expect(hostsAfter.GetMemberCount()).Should(BeZero())

		})

	})

})
