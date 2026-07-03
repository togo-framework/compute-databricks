<!-- togo-brand -->
<p align="center"><img src=".github/assets/togo-mark.svg" width="96" alt="togo" /></p>
<h1 align="center">compute-databricks</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a></sub></p>

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

MIT © fadymondy
