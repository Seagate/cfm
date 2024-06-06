// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"net/http"
	"strings"

	openapiclient "cfm/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

var _ = DescribeRegression("Host Mem Device Testing", func(sc *TestContext) {

	var (
		apiClient *openapiclient.APIClient
		hostid    string
		devid_err = "00-00"
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

	Describe("HostMemDevice", func() {

		It("should get all memory devices", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			logger.V(3).Info("get memory devices", "hostid", hostid)

			memory, httpResponse, err := apiClient.DefaultAPI.HostsGetMemoryDevices(sc.Config.Ctx, hostid).Execute()

			Expect(err).To(BeNil())
			Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
			Expect(memory).NotTo(BeNil())

			members := memory.GetMembers()

			Expect(len(members)).Should(BeNumerically("==", memory.GetMemberCount()))
		})

		It("try to get all memory devices by id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			logger.V(3).Info("try to get all memory devices by id")

			memoryList, httpResponse, err := apiClient.DefaultAPI.HostsGetMemoryDevices(sc.Config.Ctx, hostid).Execute()

			Expect(err).To(BeNil())
			Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
			Expect(memoryList).NotTo(BeNil())

			membersCnt := memoryList.GetMemberCount()

			if membersCnt != 0 {
				for _, devMember := range memoryList.GetMembers() {
					devid := strings.Split(devMember.Uri, "/")[len(strings.Split(devMember.Uri, "/"))-1]
					memory, httpResponse, err := apiClient.DefaultAPI.HostsGetMemoryDeviceById(sc.Config.Ctx, hostid, devid).Execute()

					// the device id should be valid
					Expect(err).To(BeNil())
					Expect(httpResponse).NotTo(BeNil())
					Expect(memory).NotTo(BeNil())
					Expect(memory.Id).To(ContainSubstring("memdev")) // format: memdev19-00.0
					Expect(memory.Id).To(ContainSubstring("9-00."))
				}
			} else {
				logger.V(3).Info("No device detected")
			}

		})

		It("should get error code 404 on non-exist memory device id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			logger.V(3).Info("get memory by id", "devid", devid_err)

			memory, httpResponse, err := apiClient.DefaultAPI.HostsGetMemoryDeviceById(sc.Config.Ctx, hostid, devid_err).Execute()

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
