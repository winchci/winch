package winch

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
