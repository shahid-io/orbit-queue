# 1. Initialize Go module

```go
go mod init github.com/shahid-io/orbit-queue
```

## 2. Add dependencies (e.g., gin, cron)

```go
go get github.com/gin-gonic/gin
go get github.com/robfig/cron/v3
go get github.com/sirupsen/logrus
go get github.com/prometheus/client_golang/prometheus
```

### Feature Branches

| Feature | Branch Name |
|---------|-------------|
| Initial project setup | `feature/initial-setup` |
| HTTP server + health check | `feature/api-server` |
| Core scheduler setup | `feature/scheduler-core` |
| Cron-based job execution | `feature/cron-jobs` |
| Dynamic job registration API | `feature/dynamic-jobs-api` |
| Job persistence (DB) | `feature/job-db` |
| Manual trigger support | `feature/manual-trigger` |
| Retry mechanism | `feature/job-retry` |
| Job dependencies/DAG | `feature/job-dependencies` |
| Observability & metrics (Prom) | `feature/observability` |
| Log monitoring per job | `feature/log-monitoring` |
| Dashboard UI | `feature/dashboard-ui` |
