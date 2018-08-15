package main

import (
	"net/http"

	"github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-utils/system"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	agentURL = kingpin.Flag(
		"boshdirector.agenturl", "BOSH URL ($BOSH-DIRECTOR_EXPORTER_AGENT_URL)",
	).Envar("BOSHDIRECTOR_EXPORTER_AGENT_URL").Required().String()

	agentUsername = kingpin.Flag(
		"boshdirector.agentusername", "BOSH Username ($BOSH-DIRECTOR_EXPORTER_AGENT_USERNAME)",
	).Envar("BOSHDIRECTOR_EXPORTER_AGENT_USERNAME").Required().String()

	agentPassword = kingpin.Flag(
		"boshdirector.agentpassword", "BOSH Password ($BOSH-DIRECTOR_EXPORTER_AGENT_PASSWORD)",
	).Envar("BOSHDIRECTOR_EXPORTER_AGENT_PASSWORD").Required().String()

	metricsNamespace = kingpin.Flag(
		"metrics.namespace", "Metrics Namespace ($BOSH-DIRECTOR_EXPORTER_METRICS_NAMESPACE)",
	).Envar("BOSHDIRECTOR_EXPORTER_METRICS_NAMESPACE").Default("bosh").String()

	metricsEnvironment = kingpin.Flag(
		"metrics.environment", "Environment label to be attached to metrics ($BOSH-DIRECTOR_EXPORTER_METRICS_ENVIRONMENT)",
	).Envar("BOSHDIRECTOR_EXPORTER_METRICS_ENVIRONMENT").Required().String()

	listenAddress = kingpin.Flag(
		"web.listen-address", "Address to listen on for web interface and telemetry ($BOSH-DIRECTOR_EXPORTER_WEB_LISTEN_ADDRESS)",
	).Envar("BOSHDIRECTOR_EXPORTER_WEB_LISTEN_ADDRESS").Default(":9190").String()

	metricsPath = kingpin.Flag(
		"web.telemetry-path", "Path under which to expose Prometheus metrics ($BOSH-DIRECTOR_EXPORTER_WEB_TELEMETRY_PATH)",
	).Envar("BOSHDIRECTOR_EXPORTER_WEB_TELEMETRY_PATH").Default("/metrics").String()

	authUsername = kingpin.Flag(
		"web.auth.username", "Username for web interface basic auth ($BOSH-DIRECTOR_EXPORTER_WEB_AUTH_USERNAME)",
	).Envar("BOSHDIRECTOR_EXPORTER_WEB_AUTH_USERNAME").String()

	authPassword = kingpin.Flag(
		"web.auth.password", "Password for web interface basic auth ($BOSH-DIRECTOR_EXPORTER_WEB_AUTH_PASSWORD)",
	).Envar("BOSHDIRECTOR_EXPORTER_WEB_AUTH_PASSWORD").String()

	tlsCertFile = kingpin.Flag(
		"web.tls.cert_file", "Path to a file that contains the TLS certificate (PEM format). If the certificate is signed by a certificate authority, the file should be the concatenation of the server's certificate, any intermediates, and the CA's certificate ($BOSH-DIRECTOR_EXPORTER_WEB_TLS_CERTFILE)",
	).Envar("BOSHDIRECTOR_EXPORTER_WEB_TLS_CERTFILE").ExistingFile()

	tlsKeyFile = kingpin.Flag(
		"web.tls.key_file", "Path to a file that contains the TLS private key (PEM format) ($BOSH-DIRECTOR_EXPORTER_WEB_TLS_KEYFILE)",
	).Envar("BOSH-IRECTOR_EXPORTER_WEB_TLS_KEYFILE").ExistingFile()
)

func init() {
	prometheus.MustRegister(version.NewCollector(*metricsNamespace))
}

type basicAuthHandler struct {
	handler  http.HandlerFunc
	username string
	password string
}

func (h *basicAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || username != h.username || password != h.password {
		log.Errorf("Invalid HTTP auth from `%s`", r.RemoteAddr)
		w.Header().Set("WWW-Authenticate", "Basic realm=\"metrics\"")
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	h.handler(w, r)
	return
}

func prometheusHandler() http.Handler {
	handler := prometheus.Handler()

	if *authUsername != "" && *authPassword != "" {
		handler = &basicAuthHandler{
			handler:  prometheus.Handler().ServeHTTP,
			username: *authUsername,
			password: *authPassword,
		}
	}

	return handler
}

func readCACert(CACertFile string, logger logger.Logger) (string, error) {
	if CACertFile != "" {
		fs := system.NewOsFileSystem(logger)

		CACertFileFullPath, err := fs.ExpandPath(CACertFile)
		if err != nil {
			return "", err
		}

		CACert, err := fs.ReadFileString(CACertFileFullPath)
		if err != nil {
			return "", err
		}

		return CACert, nil
	}

	return "", nil
}

func main() {
	//log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print("fbosh_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.Infoln("Starting bosh-director_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	/////
	stateClient := newAgentStateClient(*agentURL, *agentUsername, *agentPassword, true)
	stateCollector, err := newAgentStateCollector(&stateClient, *metricsEnvironment, "boshdirector")
	if err != nil {
		log.Fatalf("Could not create collector. Exitting. Error: %v", err)
	}
	prometheus.MustRegister(stateCollector)

	http.Handle(*metricsPath, prometheusHandler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>BOSH Exporter</title></head>
             <body>
             <h1>BOSH Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	if *tlsCertFile != "" && *tlsKeyFile != "" {
		log.Infoln("Listening TLS on", *listenAddress)
		log.Fatal(http.ListenAndServeTLS(*listenAddress, *tlsCertFile, *tlsKeyFile, nil))
	} else {
		log.Infoln("Listening on", *listenAddress)
		log.Fatal(http.ListenAndServe(*listenAddress, nil))
	}

}
