package main

import (
	"testing"
)

func TestGetState(t *testing.T) {
	testServer := MockServer("fakeuser", "fakepassword", mockReply)
	client := newAgentStateClient(testServer.URL, "fakeuser", "fakepassword", true)
	state, err := client.getState()
	if err != nil {
		t.Errorf("getState() returned error where no error was expected: %v", err)
	} else {
		if state.Value.AgentID != "fakeid" {
			t.Errorf("Expected state.Value.AgentID to be fakeid but got: %v", state.Value.AgentID)
		}
	}

	testServer.Close()

	testServer = MockServer("fakeuser", "fakepassword", "{}")
	client = newAgentStateClient(testServer.URL, "fakeuser", "fakepassword", true)
	_, err = client.getState()
	if err == nil {
		t.Errorf("getState() on invalid state did not return error while an error was expected")
	}

	client = newAgentStateClient(testServer.URL, "fakeuser", "wrongpassword", true)
	_, err = client.getState()
	if err == nil || err.Error() != "getState returned statuscode: 401" {
		t.Errorf("getState() with wrong creds did not return error while error was expected")
	}

	client = newAgentStateClient("wrongurl", "fakeuser", "fakepassword", true)
	_, err = client.getState()
	if err == nil {
		t.Errorf("getState() with invalid url did not return error while error was expected")
	}
	testServer.Close()
}

func TestGetStateParseFail(t *testing.T) {
	testServer := MockServer("fakeuser", "fakepassword", "{")
	defer testServer.Close()

	client := newAgentStateClient(testServer.URL, "fakeuser", "fakepassword", true)
	_, err := client.getState()

	if err == nil {
		t.Errorf("getState() on invalid data did not return parse error while error was expected")
		return
	}
}
