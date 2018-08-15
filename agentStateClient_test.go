package main

import (
	"testing"
)

func TestGetState(t *testing.T) {
	testServer := MockServer("fakeuser", "fakepassword", mockReply)
	client, _ := newAgentStateClient(testServer.URL, "fakeuser", "fakepassword", "")
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
	client, _ = newAgentStateClient(testServer.URL, "fakeuser", "fakepassword", "")
	_, err = client.getState()
	if err == nil {
		t.Errorf("getState() on invalid state did not return error while an error was expected")
	}

	client, _ = newAgentStateClient(testServer.URL, "fakeuser", "wrongpassword", "")
	_, err = client.getState()
	if err == nil || err.Error() != "getState returned statuscode: 401" {
		t.Errorf("getState() with wrong creds did not return error while error was expected")
	}

	client, _ = newAgentStateClient("wrongurl", "fakeuser", "fakepassword", "")
	_, err = client.getState()
	if err == nil {
		t.Errorf("getState() with invalid url did not return error while error was expected")
	}
	testServer.Close()
}

func TestGetStateParseFail(t *testing.T) {
	testServer := MockServer("fakeuser", "fakepassword", "{")
	defer testServer.Close()

	client, _ := newAgentStateClient(testServer.URL, "fakeuser", "fakepassword", "")
	_, err := client.getState()

	if err == nil {
		t.Errorf("getState() on invalid data did not return parse error while error was expected")
		return
	}
}
