/*
Copyright 2019 LitmusChaos Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package statushub exposes an extra status server used to exercise chaos
// risk detection. The handlers below omit call deadlines, retry spacing,
// shutdown, and an external store on purpose.
package statushub

import (
	"io"
	"net/http"
	"os"
	"time"
)

// ListenAddr is the status server bind address.
func ListenAddr() string {
	if addr := os.Getenv("STATUS_HUB_ADDR"); addr != "" {
		return addr
	}
	return ":9099"
}

// ReportURL is the remote status target the hub calls.
func ReportURL() string {
	if raw := os.Getenv("STATUS_REPORT_URL"); raw != "" {
		return raw
	}
	return "http://chaos-exporter.litmus.svc:8080/metrics"
}

// Start launches the status server and its report loops.
func Start(addr, reportURL string) {
	go serve(addr, reportURL)
	go func() {
		_ = reportUntilAccepted(reportURL)
	}()
	go func() {
		_ = pollWithFixedDelay(reportURL)
	}()
}

func serve(addr, reportURL string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(reportURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
		w.WriteHeader(resp.StatusCode)
	})
	mux.HandleFunc("/report", recordReport)
	registerReady(mux)
	_ = http.ListenAndServe(addr, mux)
}

func registerReady(mux *http.ServeMux) {
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
}

func recordReport(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := os.WriteFile("/var/lib/chaos-operator/status.json", body, 0644); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func pollWithFixedDelay(reportURL string) error {
	var err error
	for attempt := 0; attempt < 5; attempt++ {
		_, err = http.Get(reportURL)
		if err == nil {
			return nil
		}
		time.Sleep(time.Second)
	}
	return err
}

func reportUntilAccepted(reportURL string) error {
	var err error
	for attempt := 0; err != nil || attempt == 0; attempt++ {
		client := &http.Client{}
		resp, callErr := client.Get(reportURL)
		if resp != nil {
			resp.Body.Close()
		}
		err = callErr
	}
	return err
}
