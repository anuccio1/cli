package v7action_test

import (
	. "code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/actor/v7action/v7actionfakes"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Job Actions", func() {

	var (
		actor                     *Actor
		fakeCloudControllerClient *v7actionfakes.FakeCloudControllerClient

		jobURL     ccv3.JobURL
		warnings   Warnings
		err error
	)

	BeforeEach(func() {
		actor, fakeCloudControllerClient, _, _, _ = NewTestActor()
		jobURL = ccv3.JobURL("http://example.com/the-job-url")
	})

	Describe("PollJob", func() {
		When("the cc client returns success", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.PollJobReturns(ccv3.Warnings{"some-warnings"}, nil)
			})

			It("returns success and warnings", func() {
				warnings, err = actor.PollJob(jobURL)

				Expect(err).ToNot(HaveOccurred())
				Expect(warnings).To(ConsistOf("some-warnings"))
				actualJobURL := fakeCloudControllerClient.PollJobArgsForCall(0)
				Expect(actualJobURL).To(Equal(jobURL))
			})
		})

		When("the cc client returns error", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.PollJobReturns(ccv3.Warnings{"some-warnings"}, ccerror.JobFailedError{})
			})

			It("returns error and warnings", func() {
				warnings, err = actor.PollJob(jobURL)

				Expect(err).To(Equal(ccerror.JobFailedError{}))
				Expect(warnings).To(ConsistOf("some-warnings"))
				actualJobURL := fakeCloudControllerClient.PollJobArgsForCall(0)
				Expect(actualJobURL).To(Equal(jobURL))
			})
		})
	})
})
