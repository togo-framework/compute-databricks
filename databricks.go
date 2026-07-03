// Package computedatabricks is a togo compute backend that submits one-time runs
// to Databricks via the Jobs REST API (/api/2.1/jobs/runs/submit). Select with
// `togo provider:use compute databricks`.
//
//   DATABRICKS_HOST        workspace URL (https://…azuredatabricks.net)
//   DATABRICKS_TOKEN       PAT (secret)
//   DATABRICKS_CLUSTER_ID  existing cluster to run on
package computedatabricks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/togo-framework/compute"
	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("compute-databricks", togo.PriorityService+1, func(k *togo.Kernel) error {
		providers.Use(k, providers.CapCompute, "databricks", newDatabricks(k), false)
		if k.Log != nil {
			k.Log.Info("plugin active", "plugin", "compute-databricks")
		}
		return nil
	})
}

type databricks struct {
	host    string
	token   string
	cluster string
	hc      *http.Client
}

func newDatabricks(k *togo.Kernel) *databricks {
	return &databricks{
		host:    strings.TrimRight(providers.Value(k, providers.CapCompute, "databricks", "host", "", false), "/"),
		token:   providers.Value(k, providers.CapCompute, "databricks", "token", "", true),
		cluster: providers.Value(k, providers.CapCompute, "databricks", "cluster_id", "", false),
		hc:      &http.Client{Timeout: 30 * time.Second},
	}
}

// body builds the runs/submit payload for a job.
func (d *databricks) body(job compute.Job) (map[string]any, error) {
	if len(job.Cmd) == 0 {
		return nil, fmt.Errorf("compute-databricks: job %q has no Cmd (the python_file + params)", job.Name)
	}
	task := map[string]any{"python_file": job.Cmd[0], "parameters": job.Cmd[1:]}
	b := map[string]any{"run_name": job.Name, "spark_python_task": task}
	if d.cluster != "" {
		b["existing_cluster_id"] = d.cluster
	}
	return b, nil
}

func (d *databricks) Submit(ctx context.Context, job compute.Job) (compute.Run, error) {
	if d.host == "" || d.token == "" {
		return compute.Run{}, fmt.Errorf("compute-databricks: set DATABRICKS_HOST and DATABRICKS_TOKEN")
	}
	payload, err := d.body(job)
	if err != nil {
		return compute.Run{}, err
	}
	buf, _ := json.Marshal(payload)
	req, _ := http.NewRequestWithContext(ctx, "POST", d.host+"/api/2.1/jobs/runs/submit", bytes.NewReader(buf))
	req.Header.Set("Authorization", "Bearer "+d.token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := d.hc.Do(req)
	if err != nil {
		return compute.Run{}, err
	}
	defer resp.Body.Close()
	var out struct {
		RunID int64  `json:"run_id"`
		Error string `json:"error_code"`
		Msg   string `json:"message"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&out)
	if resp.StatusCode >= 300 || out.RunID == 0 {
		return compute.Run{Status: "failed", Output: out.Error + " " + out.Msg}, fmt.Errorf("databricks submit: %s %s", out.Error, out.Msg)
	}
	return compute.Run{ID: fmt.Sprint(out.RunID), Status: "running"}, nil
}
