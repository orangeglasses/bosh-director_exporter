package main

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestCollector(t *testing.T) {
	testServer := MockServer("fakeusername", "fakepassword", mockReply)
	stateClient := newAgentStateClient(testServer.URL, "fakeusername", "fakepassword", true)
	stateCollector, err := newAgentStateCollector(&stateClient, *metricsEnvironment)
	if err != nil {
		t.Errorf("newAgentStateCollector returned error where no error was expected: %v", err)
		return
	}

	descCh := make(chan *prometheus.Desc)
	go func() {
		stateCollector.Describe(descCh)
		close(descCh)
	}()

	counter := 0
	for _ = range descCh {
		counter++
	}

	if counter != 10 {
		t.Error("Did not get description for all counters")
	}

	counter = 0
	collCh := make(chan prometheus.Metric)
	go func() {
		stateCollector.Collect(collCh)
		close(collCh)
	}()

	for _ = range collCh {
		counter++
	}
	if counter < 10 {
		t.Errorf("Expected 10 metrics but received: %v", counter)
	}
}
