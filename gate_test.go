package integration_test

import (
	"net/http"

	. "."

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gate", func() {
	var (
		client         *http.Client
		correctAPIKeys = []string{
			"ExactPatternApiKey",
			"WildcardPatternApiKey",
			"SeveralPatternsApiKey",
			"IgnoringCasePatternsApiKey",
			"UniversalApiKey",
		}
	)

	BeforeEach(func() {
		client = &http.Client{}
	})

	Context("Valid API keys ", func() {
		It("Should return Bad Request when no body", func() {
			req, _ := http.NewRequest("POST", PayloadURL, nil)
			req.Header.Set(APIKeyField, correctAPIKeys[0])
			resp, _ := client.Do(req)

			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
		}, Timeout)

		DescribeTable("Authorization keys",
			func(apiKey, routingKey string, expectedCode int) {
				actualCode := SendBody(PayloadURL, APIKeyField, apiKey, routingKey).StatusCode
				Expect(actualCode).To(Equal(expectedCode))
			},
			Entry("Exact", "ExactPatternApiKey", "project.env", http.StatusOK),
			Entry("Wildcard 1", "WildcardPatternApiKey", "project.env.a", http.StatusOK),
			Entry("Wildcard 2", "WildcardPatternApiKey", "project.env.b", http.StatusOK),
			Entry("Several 1", "SeveralPatternsApiKey", "project.dev", http.StatusOK),
			Entry("Several 2", "SeveralPatternsApiKey", "project.staging", http.StatusOK),
			Entry("Several 3", "SeveralPatternsApiKey", "foo.bar", http.StatusOK),
			Entry("Ignoring Case 1", "IgnoringCasePatternsApiKey", "Project.Env", http.StatusOK),
			Entry("Ignoring Case 2", "IgnoringCasePatternsApiKey", "Foo.Bar.a", http.StatusOK),
			Entry("Ignoring Case 3", "IgnoringCasePatternsApiKey", "Foo.Bar.b", http.StatusOK),
			Entry("Ignoring Case 4", "IgnoringCasePatternsApiKey", "project.Env", http.StatusOK),
			Entry("Ignoring Case 5", "IgnoringCasePatternsApiKey", "Foo.bar.a", http.StatusOK),
			Entry("Ignoring Case 6", "IgnoringCasePatternsApiKey", "foo.Bar.b", http.StatusOK),
			Entry("Universal 1", "UniversalApiKey", "project.env", http.StatusOK),
			Entry("Universal 2", "UniversalApiKey", "project.env.a", http.StatusOK),
			Entry("Universal 3", "UniversalApiKey", "foo.bar", http.StatusOK),
			Entry("Universal 4", "UniversalApiKey", "project.Env", http.StatusOK),

			Entry("- Exact", "ExactPatternApiKey", "project.env2", http.StatusBadRequest),
			Entry("- Wildcard 1", "WildcardPatternApiKey", "project.env2.a", http.StatusBadRequest),
			Entry("- Wildcard 2", "WildcardPatternApiKey", "project.env2.b", http.StatusBadRequest),
			Entry("- Several 1", "SeveralPatternsApiKey", "project.dev2", http.StatusBadRequest),
			Entry("- Several 2", "SeveralPatternsApiKey", "project2.staging", http.StatusBadRequest),
			Entry("- Several 3", "SeveralPatternsApiKey", "foo.bar2", http.StatusBadRequest),
			Entry("- Ignoring 1", "IgnoringCasePatternsApiKey", "Project2.Env", http.StatusBadRequest),
			Entry("- Ignoring 2", "IgnoringCasePatternsApiKey", "Foo.Bar2.a", http.StatusBadRequest),
			Entry("- Ignoring 3", "IgnoringCasePatternsApiKey", "Foo2.Bar.b", http.StatusBadRequest),
			Entry("- Ignoring 4", "IgnoringCasePatternsApiKey", "project2.Env", http.StatusBadRequest),
			Entry("- Ignoring 5", "IgnoringCasePatternsApiKey", "Foo2.bar2.a", http.StatusBadRequest),
			Entry("- Ignoring 6", "IgnoringCasePatternsApiKey", "foo.Bar2.b", http.StatusBadRequest),
		)
	})

	Context("Wrong API keys ", func() {
		It("Should return Unauthorized when no apikey", func() {
			resp, _ := http.Post(PayloadURL, "multipart/form-data", nil)
			Expect(resp.StatusCode).To(Equal(http.StatusUnauthorized))
		})
	})
})
