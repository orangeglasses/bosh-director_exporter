package main

import (
	"fmt"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type metric struct {
	gauge prometheus.Gauge
	value float64
}

type agentStateCollector struct {
	client  *agentStateClient
	metrics map[string]metric
}

//wrap float conversion in stuct to simplify error handling
type toFloat struct {
	Err error
}

func (t toFloat) do(input string) (result float64) {
	result = 0.0

	if t.Err != nil {
		return
	}

	result, t.Err = strconv.ParseFloat(input, 64)
	return
}

////

func (c *agentStateCollector) getVitalsMetrics() error {
	state, err := c.client.getState()
	if err != nil {
		return fmt.Errorf("Cannot get director state: %v", err.Error())
	}

	vitals := state.Value.Vitals
	if vitals == nil {
		return fmt.Errorf("No vitals received from getState()")
	}

	values := make(map[string]float64)

	var floatConverter toFloat
	values["CPUSys"] = floatConverter.do(vitals.CPU.Sys)
	values["CPUUser"] = floatConverter.do(vitals.CPU.User)
	values["CPUWait"] = floatConverter.do(vitals.CPU.Wait)

	values["DiskPersistent"] = floatConverter.do(vitals.Disk["persistent"].Percent)
	values["DiskEphemeral"] = floatConverter.do(vitals.Disk["ephemeral"].Percent)
	values["DiskSystem"] = floatConverter.do(vitals.Disk["system"].Percent)

	values["MemPercent"] = floatConverter.do(vitals.Mem.Percent)
	values["MemKb"] = floatConverter.do(vitals.Mem.Kb)

	values["SwapPercent"] = floatConverter.do(vitals.Swap.Percent)
	values["SwapKb"] = floatConverter.do(vitals.Swap.Kb)

	if floatConverter.Err != nil {
		return err
	}

	for k, v := range values {
		m := c.metrics[k]
		m.value = v
		c.metrics[k] = m
	}

	return nil
}

func (c *agentStateCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, v := range c.metrics {
		v.gauge.Describe(ch)
	}

}

func (c *agentStateCollector) Collect(ch chan<- prometheus.Metric) {
	err := c.getVitalsMetrics()
	if err != nil {
		log.Errorf("No metrics received: %v", err)
		defer close(ch)
		return
	}
	for _, v := range c.metrics {
		ch <- prometheus.MustNewConstMetric(v.gauge.Desc(), prometheus.GaugeValue, v.value)
	}
}

func newAgentStateCollector(stateClient *agentStateClient, environment, namespace string) (*agentStateCollector, error) {
	state, err := stateClient.getState()
	if err != nil {
		log.Errorf("Cannot get agent state: %v", err.Error())
		return nil, err
	}

	log.Infof("Scraping Director: %v", state.Value.Deployment)

	constLabels := prometheus.Labels{"environment": environment,
		"bosh_name": state.Value.Deployment}

	collector := &agentStateCollector{
		client:  stateClient,
		metrics: make(map[string]metric),
	}

	collector.metrics["CPUSys"] = metric{
		prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   "vitals",
				Name:        "cpu_sys",
				Help:        "CPU SYS Usage",
				ConstLabels: constLabels,
			},
		), 0}

	collector.metrics["CPUUser"] = metric{
		prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   "vitals",
				Name:        "cpu_user",
				Help:        "CPU USER Usage",
				ConstLabels: constLabels,
			},
		), 0}

	collector.metrics["CPUWait"] = metric{
		prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   "vitals",
				Name:        "cpu_wait",
				Help:        "CPU WAIT Usage",
				ConstLabels: constLabels,
			},
		), 0}

	collector.metrics["DiskPersistent"] = metric{
		prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   "vitals",
				Name:        "disk_persistent_percentage",
				Help:        "Percentage Persistent Disk Usage",
				ConstLabels: constLabels,
			},
		), 0}

	collector.metrics["DiskEphemeral"] = metric{
		prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   "vitals",
				Name:        "disk_ephemeral_percentage",
				Help:        "Percentage ephemeral Disk Usage",
				ConstLabels: constLabels,
			},
		), 0}

	collector.metrics["DiskSystem"] = metric{
		prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   "vitals",
				Name:        "disk_system_percentage",
				Help:        "Percentage System Disk Usage",
				ConstLabels: constLabels,
			},
		), 0}

	collector.metrics["MemPercent"] = metric{
		prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   "vitals",
				Name:        "mem_percentage",
				Help:        "Percentage RAM Usage",
				ConstLabels: constLabels,
			},
		), 0}

	collector.metrics["MemKb"] = metric{
		prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   "vitals",
				Name:        "mem_kb",
				Help:        "KiloBytes RAM Usage",
				ConstLabels: constLabels,
			},
		), 0}

	collector.metrics["SwapPercent"] = metric{
		prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   "vitals",
				Name:        "swap_percentage",
				Help:        "Percentage Swap Usage",
				ConstLabels: constLabels,
			},
		), 0}
	collector.metrics["SwapKb"] = metric{
		prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   "vitals",
				Name:        "swap_kb",
				Help:        "KiloBytes Swap Usage",
				ConstLabels: constLabels,
			},
		), 0}

	return collector, nil
}
