// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotalservices/goulash/handler"
	"github.com/pivotalservices/slack"
)

type FakeSlackAPI struct {
	PostMessageStub        func(channelID string, text string, params slack.PostMessageParameters) (channel string, timestamp string, err error)
	postMessageMutex       sync.RWMutex
	postMessageArgsForCall []struct {
		channelID string
		text      string
		params    slack.PostMessageParameters
	}
	postMessageReturns struct {
		result1 string
		result2 string
		result3 error
	}
	InviteGuestStub        func(teamname, channelID, firstName, lastName, emailAddress string) error
	inviteGuestMutex       sync.RWMutex
	inviteGuestArgsForCall []struct {
		teamname     string
		channelID    string
		firstName    string
		lastName     string
		emailAddress string
	}
	inviteGuestReturns struct {
		result1 error
	}
	InviteRestrictedStub        func(teamname, channelID, firstName, lastName, emailAddress string) error
	inviteRestrictedMutex       sync.RWMutex
	inviteRestrictedArgsForCall []struct {
		teamname     string
		channelID    string
		firstName    string
		lastName     string
		emailAddress string
	}
	inviteRestrictedReturns struct {
		result1 error
	}
	GetGroupsStub        func(excludeArchived bool) ([]slack.Group, error)
	getGroupsMutex       sync.RWMutex
	getGroupsArgsForCall []struct {
		excludeArchived bool
	}
	getGroupsReturns struct {
		result1 []slack.Group
		result2 error
	}
}

func (fake *FakeSlackAPI) PostMessage(channelID string, text string, params slack.PostMessageParameters) (channel string, timestamp string, err error) {
	fake.postMessageMutex.Lock()
	fake.postMessageArgsForCall = append(fake.postMessageArgsForCall, struct {
		channelID string
		text      string
		params    slack.PostMessageParameters
	}{channelID, text, params})
	fake.postMessageMutex.Unlock()
	if fake.PostMessageStub != nil {
		return fake.PostMessageStub(channelID, text, params)
	} else {
		return fake.postMessageReturns.result1, fake.postMessageReturns.result2, fake.postMessageReturns.result3
	}
}

func (fake *FakeSlackAPI) PostMessageCallCount() int {
	fake.postMessageMutex.RLock()
	defer fake.postMessageMutex.RUnlock()
	return len(fake.postMessageArgsForCall)
}

func (fake *FakeSlackAPI) PostMessageArgsForCall(i int) (string, string, slack.PostMessageParameters) {
	fake.postMessageMutex.RLock()
	defer fake.postMessageMutex.RUnlock()
	return fake.postMessageArgsForCall[i].channelID, fake.postMessageArgsForCall[i].text, fake.postMessageArgsForCall[i].params
}

func (fake *FakeSlackAPI) PostMessageReturns(result1 string, result2 string, result3 error) {
	fake.PostMessageStub = nil
	fake.postMessageReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSlackAPI) InviteGuest(teamname string, channelID string, firstName string, lastName string, emailAddress string) error {
	fake.inviteGuestMutex.Lock()
	fake.inviteGuestArgsForCall = append(fake.inviteGuestArgsForCall, struct {
		teamname     string
		channelID    string
		firstName    string
		lastName     string
		emailAddress string
	}{teamname, channelID, firstName, lastName, emailAddress})
	fake.inviteGuestMutex.Unlock()
	if fake.InviteGuestStub != nil {
		return fake.InviteGuestStub(teamname, channelID, firstName, lastName, emailAddress)
	} else {
		return fake.inviteGuestReturns.result1
	}
}

func (fake *FakeSlackAPI) InviteGuestCallCount() int {
	fake.inviteGuestMutex.RLock()
	defer fake.inviteGuestMutex.RUnlock()
	return len(fake.inviteGuestArgsForCall)
}

func (fake *FakeSlackAPI) InviteGuestArgsForCall(i int) (string, string, string, string, string) {
	fake.inviteGuestMutex.RLock()
	defer fake.inviteGuestMutex.RUnlock()
	return fake.inviteGuestArgsForCall[i].teamname, fake.inviteGuestArgsForCall[i].channelID, fake.inviteGuestArgsForCall[i].firstName, fake.inviteGuestArgsForCall[i].lastName, fake.inviteGuestArgsForCall[i].emailAddress
}

func (fake *FakeSlackAPI) InviteGuestReturns(result1 error) {
	fake.InviteGuestStub = nil
	fake.inviteGuestReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSlackAPI) InviteRestricted(teamname string, channelID string, firstName string, lastName string, emailAddress string) error {
	fake.inviteRestrictedMutex.Lock()
	fake.inviteRestrictedArgsForCall = append(fake.inviteRestrictedArgsForCall, struct {
		teamname     string
		channelID    string
		firstName    string
		lastName     string
		emailAddress string
	}{teamname, channelID, firstName, lastName, emailAddress})
	fake.inviteRestrictedMutex.Unlock()
	if fake.InviteRestrictedStub != nil {
		return fake.InviteRestrictedStub(teamname, channelID, firstName, lastName, emailAddress)
	} else {
		return fake.inviteRestrictedReturns.result1
	}
}

func (fake *FakeSlackAPI) InviteRestrictedCallCount() int {
	fake.inviteRestrictedMutex.RLock()
	defer fake.inviteRestrictedMutex.RUnlock()
	return len(fake.inviteRestrictedArgsForCall)
}

func (fake *FakeSlackAPI) InviteRestrictedArgsForCall(i int) (string, string, string, string, string) {
	fake.inviteRestrictedMutex.RLock()
	defer fake.inviteRestrictedMutex.RUnlock()
	return fake.inviteRestrictedArgsForCall[i].teamname, fake.inviteRestrictedArgsForCall[i].channelID, fake.inviteRestrictedArgsForCall[i].firstName, fake.inviteRestrictedArgsForCall[i].lastName, fake.inviteRestrictedArgsForCall[i].emailAddress
}

func (fake *FakeSlackAPI) InviteRestrictedReturns(result1 error) {
	fake.InviteRestrictedStub = nil
	fake.inviteRestrictedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSlackAPI) GetGroups(excludeArchived bool) ([]slack.Group, error) {
	fake.getGroupsMutex.Lock()
	fake.getGroupsArgsForCall = append(fake.getGroupsArgsForCall, struct {
		excludeArchived bool
	}{excludeArchived})
	fake.getGroupsMutex.Unlock()
	if fake.GetGroupsStub != nil {
		return fake.GetGroupsStub(excludeArchived)
	} else {
		return fake.getGroupsReturns.result1, fake.getGroupsReturns.result2
	}
}

func (fake *FakeSlackAPI) GetGroupsCallCount() int {
	fake.getGroupsMutex.RLock()
	defer fake.getGroupsMutex.RUnlock()
	return len(fake.getGroupsArgsForCall)
}

func (fake *FakeSlackAPI) GetGroupsArgsForCall(i int) bool {
	fake.getGroupsMutex.RLock()
	defer fake.getGroupsMutex.RUnlock()
	return fake.getGroupsArgsForCall[i].excludeArchived
}

func (fake *FakeSlackAPI) GetGroupsReturns(result1 []slack.Group, result2 error) {
	fake.GetGroupsStub = nil
	fake.getGroupsReturns = struct {
		result1 []slack.Group
		result2 error
	}{result1, result2}
}

var _ handler.SlackAPI = new(FakeSlackAPI)
