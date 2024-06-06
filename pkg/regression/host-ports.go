// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"fmt"
	"net/http"
	"strings"

	openapiclient "cfm/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

var _ = DescribeRegression("Host Ports Testing", func(sc *TestContext) {

	var (
		apiClient *openapiclient.APIClient
		hostid    string
		// portid    = "port19-00"
		badPortid = "portX"
	)

	BeforeEach(func() {

		By("setup client for testing")

		apiClient = openapiclient.NewAPIClient(sc.Config.Config)

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
		Expect(err).To(BeNil())
		Expect(httpResponse.StatusCode).To(Equal(http.StatusCreated))
		Expect(host).NotTo(BeNil())

		if httpResponse.StatusCode == http.StatusCreated && err == nil {
			hostid = host.GetId()
		}

	})

	Describe("HostPorts", func() {

		It("should get all ports", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			logger.V(3).Info("get ports", "hostid", hostid)

			ports, httpResponse, err := apiClient.DefaultAPI.HostsGetPorts(sc.Config.Ctx, hostid).Execute()

			Expect(err).To(BeNil())
			Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
			Expect(ports).NotTo(BeNil())

			members := ports.GetMembers()
			Expect(len(members)).Should(BeNumerically("==", ports.GetMemberCount()))
			uriExpected := fmt.Sprintf("/cfm/v1/hosts/%s/ports/", hostid)
			for _, member := range members {
				uri := member.GetUri()
				Expect(uri).To(ContainSubstring(uriExpected))
				tokens := strings.Split(uri, "/")
				Expect(len(tokens)).Should(BeNumerically(">", 0))
				id := tokens[len(tokens)-1]
				Expect(id).To(ContainSubstring("port"))
				Expect(id).To(ContainSubstring("9-00"))
			}
		})

		It("should get error code 404 on non-existent port id", func() {

			logger := klog.FromContext(sc.Config.Ctx)

			logger.V(3).Info("get port by id", "badPortid", badPortid)

			port, httpResponse, err := apiClient.DefaultAPI.HostsGetPortById(sc.Config.Ctx, hostid, badPortid).Execute()

			Expect(err).NotTo(BeNil())
			Expect(httpResponse).NotTo(BeNil())
			Expect(httpResponse.StatusCode).To(Equal(http.StatusNotFound))
			Expect(port).To(BeNil())
		})

		//TODO: Add more tests after real (or emulated) cxl connections exist between host and blade
	})

	AfterEach(func() {

		By("clean up test client")

		apiClient = openapiclient.NewAPIClient(sc.Config.Config)

		request := apiClient.DefaultAPI.HostsDeleteById(sc.Config.Ctx, hostid)

		_, _, err := request.Execute()

		Expect(err).To(BeNil())

	})

})
