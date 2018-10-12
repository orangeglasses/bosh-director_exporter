# bosh-director_exporter
Prometheus exporter for BOSH Director metrics
This exporter queries the BOSh agent running on the BOSH director, extracts the vitals and exposes this as prometheus scrapable information

building the binary: go build -o ./bosh-director_exporter

Commandline arguments / environment variables:
- boshdirector.agenturl / BOSHDIRECTOR_EXPORTER_AGENT_URL (https://<director ip>:6868/agent)
- boshdirector.agentusername / BOSHDIRECTOR_EXPORTER_AGENT_USERNAME (For Pivotal PAS: vcap, for opensource: mbus)
- boshdirector.agentpassword / BOSHDIRECTOR_EXPORTER_AGENT_PASSWORD (Password for use specified above)
- boshdirector.agent-ca-file / OSHDIRECTOR_EXPORTER_AGENT_CA_FILE (director CA cert)
- metrics.namespace / BOSHDIRECTOR_EXPORTER_METRICS_NAMESPACE (prometheus namespace)
- metrics.environment / BOSHDIRECTOR_EXPORTER_METRICS_ENVIRONMENT (prometheus environment)
- web.listen-address / BOSHDIRECTOR_EXPORTER_WEB_LISTEN_ADDRESS (default to :9191)
- web.telemetry-path / BOSHDIRECTOR_EXPORTER_WEB_TELEMETRY_PATH (defaults to /metrics)
- web.auth.username / BOSHDIRECTOR_EXPORTER_WEB_AUTH_USERNAME (optional username for metrics endpoint)
- web.auth.password / BOSHDIRECTOR_EXPORTER_WEB_AUTH_PASSWORD
- web.tls.cert_file / BOSHDIRECTOR_EXPORTER_WEB_TLS_CERTFILE (Path to TLS cert in PEM format)
- web.tls.key_file / BOSHDIRECTOR_EXPORTER_WEB_TLS_KEYFILE (Path to private key in PEM format)
