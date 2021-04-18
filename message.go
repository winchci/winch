/*
winch - Universal Build and Release Tool
Copyright (C) 2021 Ketch Kloud, Inc.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
*/

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
