package slackapi_test

import (
	"github.com/pivotalservices/slack"

	"github.com/pivotalservices/goulash/slackapi"

	fakeslackapi "github.com/pivotalservices/goulash/slackapi/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Channel", func() {
	Describe("Name", func() {
		var fakeSlackAPI *fakeslackapi.FakeSlackAPI
		var channel slackapi.Channel

		BeforeEach(func() {
			fakeSlackAPI = &fakeslackapi.FakeSlackAPI{}
		})

		Describe("when the group's name is slackapi.PrivateGroupName", func() {
			BeforeEach(func() {
				channel = slackapi.NewChannel(slackapi.PrivateGroupName, "C1234")
			})

			It("tries to find the group's name in Slack, excluding archived groups", func() {
				channel.Name(fakeSlackAPI)
				Ω(fakeSlackAPI.GetGroupsCallCount()).To(Equal(1))

				excludeArchived := fakeSlackAPI.GetGroupsArgsForCall(0)
				Ω(excludeArchived).To(BeTrue())
			})

			It("only hits the Slack API once", func() {
				channel.Name(fakeSlackAPI)
				channel.Name(fakeSlackAPI)
				Ω(fakeSlackAPI.GetGroupsCallCount()).To(Equal(1))
			})

			Describe("when the group is found in Slack", func() {
				BeforeEach(func() {
					fakeSlackAPI.GetGroupsReturns([]slack.Group{
						{Name: "channel-name", BaseChannel: slack.BaseChannel{Id: "C1234"}},
					}, nil)
				})

				It("returns the name associated with the found group", func() {
					Ω(channel.Name(fakeSlackAPI)).To(Equal("channel-name"))
				})
			})

			Describe("when the group's name is not found in Slack", func() {
				BeforeEach(func() {
					fakeSlackAPI.GetGroupsReturns([]slack.Group{
						{Name: "other-channel-name", BaseChannel: slack.BaseChannel{Id: "C9999"}},
					}, nil)
				})

				It("returns slackapi.PrivateGroupName", func() {
					Ω(channel.Name(fakeSlackAPI)).To(Equal(slackapi.PrivateGroupName))
				})
			})
		})

		Describe("when the group's name is not slackapi.PrivateGroupName", func() {
			BeforeEach(func() {
				channel = slackapi.NewChannel("channel-name", "C1234")
			})

			It("returns the name", func() {
				Ω(channel.Name(fakeSlackAPI)).To(Equal("channel-name"))
			})

			It("does not call Slack to find the group's name", func() {
				channel.Name(fakeSlackAPI)
				Ω(fakeSlackAPI.GetGroupsCallCount()).To(Equal(0))
			})
		})
	})
})