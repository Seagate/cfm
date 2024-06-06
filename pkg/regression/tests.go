// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	. "github.com/onsi/ginkgo/v2"
)

type test struct {
	text string
	body func(*TestContext)
}

var tests []test

// DescribeRegression must be used instead of the usual Ginkgo Describe to
// register a test block. The difference is that the body function
// will be called multiple times with the right context (when
// setting up a Ginkgo suite or a testing.T test, with the right
// configuration).
func DescribeRegression(text string, body func(*TestContext)) bool {
	tests = append(tests, test{text, body})
	return true
}

// registerTestsInGinkgo invokes the actual Gingko Describe
// for the tests registered earlier with DescribeSanity.
func registerTestsInGinkgo(sc *TestContext) {
	for _, test := range tests {
		test := test
		Describe(test.text, func() {
			BeforeEach(func() {
				sc.Setup()
			})

			test.body(sc)

			AfterEach(func() {
				sc.Teardown()
			})
		})
	}
}
