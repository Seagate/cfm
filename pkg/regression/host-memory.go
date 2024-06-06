// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"net/http"

	openapiclient "cfm/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

var _ = DescribeRegression("Host Memory Testing", func(sc *TestContext) {

	var (
		apiClient    *openapiclient.APIClient
		hostid       string
		memoryid     = "node0"
		memoryid_err = "node99 "
	)

	BeforeEach(func() {

		By("setup client for testing")

		apiClient = openapiclient.NewAPIClient(sc.Config.Config)

		// add a host before testing
		request := apiClient.DefaultAPI.HostsPost(sc.Config.Ctx)

		credentials := openapiclient.Credentials{
			Username:  sc.Config.CxlHostEndpoints[0].Username,
			Password:  sc.Config.CxlHostEndpoints[0].Password,
			IpAddress: sc.Config.CxlHostEndpoints[0].IpAddress,
			Port:      sc.Config.CxlHostEndpoints[0].Port,
			Insecure:  &sc.Config.CxlHostEndpoints[0].Insecure,
			Protocol:  &sc.Config.CxlHostEndpoints[0].Protocol,
		}

		request = request.Credentials(credentials)

		host, httpResponse, err := request.Execute()

		if httpResponse.StatusCode == http.StatusCreated && err == nil {
			hostid = host.GetId()
		}

	})

	Describe("HostMemory", func() {

		It("should get all memory", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			logger.V(3).Info("get memory", "hostid", hostid)

			memory, httpResponse, err := apiClient.DefaultAPI.HostGetMemory(sc.Config.Ctx, hostid).Execute()

			Expect(err).To(BeNil())
			Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
			Expect(memory).NotTo(BeNil())

			members := memory.GetMembers()

			Expect(len(members)).Should(BeNumerically("==", memory.GetMemberCount()))
		})

		It("should get a specific memory by id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			logger.V(3).Info("get memory by id", "memoryid", memoryid)

			memory, httpResponse, err := apiClient.DefaultAPI.HostsGetMemoryById(sc.Config.Ctx, hostid, memoryid).Execute()

			// the memoryid is always valid
			Expect(err).To(BeNil())
			Expect(httpResponse).NotTo(BeNil())
			Expect(memory).NotTo(BeNil())

		})

		It("should get error code 404 on non-exist memory id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			logger.V(3).Info("get memory by id", "memoryid", memoryid_err)

			memory, httpResponse, err := apiClient.DefaultAPI.HostsGetMemoryById(sc.Config.Ctx, hostid, memoryid_err).Execute()

			// the memoryid is invalid, so error will happen here and the specific memory and http response are empty
			Expect(err).NotTo(BeNil())
			Expect(httpResponse).NotTo(BeNil())
			Expect(httpResponse.StatusCode).To(Equal(http.StatusNotFound))
			Expect(memory).To(BeNil())

		})

	})

	AfterEach(func() {

		By("clean up test client")

		apiClient = openapiclient.NewAPIClient(sc.Config.Config)

		request := apiClient.DefaultAPI.HostsDeleteById(sc.Config.Ctx, hostid)

		_, _, err := request.Execute()

		Expect(err).To(BeNil())

	})

})
