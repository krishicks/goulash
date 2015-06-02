package action

import (
	"strings"

	"github.com/pivotal-golang/lager"
	"github.com/pivotalservices/goulash/slackapi"
)

var (
	uninvitableUserNotFoundMessageFmt = "There is no user here with the email address '%s'. %s"
	userInfoMessageFmt                = "%s %s (%s) is a Slack %s, with the username <@%s>."
	userNotFoundMessageFmt            = "There is no user here with the email address '%s'. You can invite them to Slack as a guest or a restricted account. Type `/butler help` for more information."

	membershipFull               = "full member"
	membershipRestrictedAccount  = "restricted account"
	membershipSingleChannelGuest = "single-channel guest"
)

// Action represents an action that is able to be performed by the server.
type Action interface {
	Do() (string, error)
}

// GuardedAction is an Action with prerequisites described in Check().
type GuardedAction interface {
	Check() error
}

// AuditableAction is an Action that should have an audit log entry created.
type AuditableAction interface {
	AuditMessage() string
}

// New creates a new Action based on the command provided.
func New(
	channel slackapi.Channel,
	commanderName string,
	commanderID string,
	text string,

	api slackapi.SlackAPI,
	slackTeamName string,
	slackUserID string,
	uninvitableDomain string,
	uninvitableMessage string,
	logger lager.Logger,
) Action {
	command, commandParams := commandAndParams(text)

	switch command {
	case "help":
		return help{}

	case "info":
		return userInfo{
			params: commandParams,

			api:                api,
			requestingUser:     commanderName,
			slackTeamName:      slackTeamName,
			uninvitableDomain:  uninvitableDomain,
			uninvitableMessage: uninvitableMessage,
			logger:             logger,
		}

	case "invite-guest":
		return inviteGuest{
			params: commandParams,

			api:                api,
			channel:            channel,
			invitingUser:       commanderName,
			slackTeamName:      slackTeamName,
			slackUserID:        slackUserID,
			uninvitableDomain:  uninvitableDomain,
			uninvitableMessage: uninvitableMessage,
			logger:             logger,
		}

	case "invite-restricted":
		return inviteRestricted{
			params: commandParams,

			api:                api,
			channel:            channel,
			invitingUser:       commanderName,
			slackTeamName:      slackTeamName,
			slackUserID:        slackUserID,
			uninvitableDomain:  uninvitableDomain,
			uninvitableMessage: uninvitableMessage,
			logger:             logger,
		}

	default:
		return help{}
	}
}

func commandAndParams(text string) (string, []string) {
	var command string
	var commandParams []string

	if commandSep := strings.IndexByte(text, 0x20); commandSep > 0 {
		command = text[:commandSep]
		commandParams = strings.Split(text[commandSep+1:], " ")
	} else {
		command = text
	}

	return command, commandParams
}

func uninvitableEmail(emailAddress string, uninvitableDomain string) bool {
	return len(uninvitableDomain) > 0 && strings.HasSuffix(emailAddress, uninvitableDomain)
}
