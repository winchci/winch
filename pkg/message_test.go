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

package pkg

import "testing"

func TestParseMessage(t *testing.T) {
	m := ParseMessage("fix!(tls): enable TLS by default\n\nbody\n\nmore body\n\nBREAKING CHANGE: something is broken now")
	if !m.IsBreaking {
		t.Fail()
	}
	if m.Type != NewType("fix") {
		t.Fail()
	}
	if m.Subject != "enable TLS by default" {
		t.Fail()
	}
	if m.Scope != "tls" {
		t.Fail()
	}
	if m.BreakingMessage != "something is broken now" {
		t.Fail()
	}
}

func TestCarriageReturn(t *testing.T) {
	m := ParseMessage("User WIP (#18)\n\n* ses wip\r\n\r\n* ses wip\r\n\r\n* User wip")
	if m.Subject != "User WIP (#18)" {
		t.Fail()
	}
}
