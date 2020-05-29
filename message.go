package winch

import (
	"fmt"
	"regexp"
	"strings"
)

type Message struct {
	Type            CommitType
	IsBreaking      bool
	Scope           string
	Subject         string
	Body            string
	BreakingMessage string
}

func (m Message) Title() string {
	return m.Type.Title()
}

func (m Message) ScopePrefix() string {
	if len(m.Scope) > 0 {
		return fmt.Sprintf("%s: ", m.Scope)
	}

	return ""
}

var (
	messageRE            = regexp.MustCompile(`(?:(?P<type>[a-z_]+)(?P<breaking>!)?(\((?P<scope>[^)]+)\))?: )?(?P<subject>[^\r\n]+)(?P<body>(?s:.*))`)
	breakingChangeMarker = "BREAKING CHANGE: "
	defaultType          = NewType("fix")
)

func parseBody(body string) (string, string) {
	i := strings.Index(body, breakingChangeMarker)

	var breakingMsg string
	if i > -1 {
		breakingMsg = strings.TrimSpace(body[(i + len(breakingChangeMarker)):])
		body = strings.TrimSpace(body[0:i])
	}

	return body, breakingMsg
}

func ParseMessage(message string) *Message {
	messageRE.Longest()

	match := messageRE.FindAllStringSubmatch(message, -1)
	if len(match) < 1 || len(match[0]) < 7 {
		subject, breakingMsg := parseBody(strings.TrimSpace(message))

		return &Message{
			Type:            defaultType,
			IsBreaking:      false,
			Scope:           "",
			Subject:         subject,
			Body:            "",
			BreakingMessage: breakingMsg,
		}
	}

	body, breakingMsg := parseBody(strings.TrimSpace(match[0][6]))

	return &Message{
		Type:            NewType(match[0][1]),
		IsBreaking:      match[0][2] == "!" || len(breakingMsg) > 0,
		Scope:           match[0][4],
		Subject:         strings.TrimSpace(match[0][5]),
		Body:            body,
		BreakingMessage: breakingMsg,
	}
}
