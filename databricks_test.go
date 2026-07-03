package computedatabricks

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/togo-framework/compute"
)

func TestSubmitPostsRunAndParsesRunID(t *testing.T) {
	var gotPath, gotAuth string
	var gotBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		_, _ = w.Write([]byte(`{"run_id": 4242}`))
	}))
	defer srv.Close()

	d := &databricks{host: srv.URL, token: "tok", cluster: "c-1", hc: srv.Client()}
	run, err := d.Submit(context.Background(), compute.Job{Name: "etl", Cmd: []string{"main.py", "--date=today"}})
	if err != nil {
		t.Fatalf("submit: %v", err)
	}
	if run.ID != "4242" || run.Status != "running" {
		t.Fatalf("run = %+v", run)
	}
	if gotPath != "/api/2.1/jobs/runs/submit" || gotAuth != "Bearer tok" {
		t.Fatalf("bad request: path=%s auth=%s", gotPath, gotAuth)
	}
	if gotBody["existing_cluster_id"] != "c-1" || gotBody["spark_python_task"] == nil {
		t.Fatalf("bad body: %v", gotBody)
	}
}

func TestSubmitRequiresConfig(t *testing.T) {
	if _, err := (&databricks{}).Submit(context.Background(), compute.Job{Cmd: []string{"x"}}); err == nil {
		t.Fatal("expected error without host/token")
	}
}
