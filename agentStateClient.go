package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	boshaction "github.com/cloudfoundry/bosh-agent/agent/action"
)

type agentState struct {
	Value boshaction.GetStateV1ApplySpec `json:"value"`
}

type agentStateClient struct {
	url        string
	username   string
	password   string
	httpClient *http.Client
}

func newAgentStateClient(url, username, password, certfile string) (agentStateClient, error) {
	caCertPool := x509.NewCertPool()

	insecure := true
	if certfile != "" {
		insecure = false

		caCert, err := ioutil.ReadFile(certfile)
		if err != nil {
			return agentStateClient{}, fmt.Errorf("Unable to load cert file: %v", err)
		}

		caCertPool.AppendCertsFromPEM(caCert)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecure,
			RootCAs:            caCertPool,
		},
	}

	httpClient := &http.Client{
		Transport: tr,
		Timeout:   20 * time.Second,
	}
	return agentStateClient{url: url,
		username:   username,
		password:   password,
		httpClient: httpClient,
	}, nil
}

func (a agentStateClient) getState() (agentState, error) {
	request, _ := http.NewRequest(http.MethodPost, a.url, strings.NewReader("{\"method\":\"get_state\",\"arguments\":[\"full\"], \"reply_to\": \"Bosh::Agent::HTTPClient\"}"))
	request.SetBasicAuth(a.username, a.password)

	resp, err := a.httpClient.Do(request)
	if err != nil {
		return agentState{}, err
	}

	if resp.StatusCode != 200 {
		return agentState{}, fmt.Errorf("getState returned statuscode: %v", resp.StatusCode)
	}

	var state agentState
	err = json.NewDecoder(resp.Body).Decode(&state)
	if err != nil {
		return agentState{}, err
	}

	if state.Value.AgentID == "" {
		return agentState{}, fmt.Errorf("No valid state received")
	}

	defer resp.Body.Close()
	return state, nil
}
