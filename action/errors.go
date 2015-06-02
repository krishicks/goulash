package action

import "fmt"

var (
	uninvitableDomainErrFmt = "Users for the '%s' domain are unable to be invited through /butler. %s"
)

// ChannelNotVisibleErr is an error.
type ChannelNotVisibleErr struct {
	slackUserID string
}

// NewChannelNotVisibleErr returns a new ChannelNotVisibleErr.
func NewChannelNotVisibleErr(slackUserID string) ChannelNotVisibleErr {
	return ChannelNotVisibleErr{
		slackUserID: slackUserID,
	}
}

func (e ChannelNotVisibleErr) Error() string {
	return fmt.Sprintf(channelNotVisibleErrFmt, e.slackUserID, e.slackUserID, e.slackUserID, e.slackUserID)
}

// UninvitableDomainErr is an error.
type UninvitableDomainErr struct {
	uninvitableDomain  string
	uninvitableMessage string
}

// NewUninvitableDomainErr returns a new UninvitableDomainErr.
func NewUninvitableDomainErr(uninvitableDomain string, uninvitableMessage string) UninvitableDomainErr {
	return UninvitableDomainErr{
		uninvitableDomain:  uninvitableDomain,
		uninvitableMessage: uninvitableMessage,
	}
}

func (e UninvitableDomainErr) Error() string {
	return fmt.Sprintf(uninvitableDomainErrFmt, e.uninvitableDomain, e.uninvitableMessage)
}
