<!-- togo-header -->
<div align="center">
  <picture><source media="(prefers-color-scheme: dark)" srcset=".github/assets/togo-mark-dark.svg" /><img src=".github/assets/togo-mark.svg" alt="ToGO" height="64" /></picture>
  <h1>togo-framework/compute-databricks</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1F8A99" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/compute-databricks"><img src="https://pkg.go.dev/badge/github.com/togo-framework/compute-databricks.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/compute-databricks
```

<!-- /togo-header -->

A togo **compute** backend that submits one-time runs to **Databricks** via the
Jobs REST API (`/api/2.1/jobs/runs/submit`).

```bash
togo install togo-framework/compute-databricks
togo provider:use compute databricks
togo config:set DATABRICKS_HOST https://xxx.cloud.databricks.com
togo config:set DATABRICKS_CLUSTER_ID 0101-… 
# DATABRICKS_TOKEN in .env (secret)
```

| Key | Purpose |
|---|---|
| `DATABRICKS_HOST` | workspace URL |
| `DATABRICKS_TOKEN` | PAT (secret) |
| `DATABRICKS_CLUSTER_ID` | existing cluster |
<!-- togo-sponsors -->
---

<div align="center">
  <h3>💎 Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
    &nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
