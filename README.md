# Orbit-Queue 🚀

Orbit-Queue is a **reliable, flexible, and observable job/task scheduling system** built with Go. It supports cron-based execution, dynamic job registration, retries, job dependencies, and system observability — ideal for robust and scalable job orchestration.

---

## ✨ Features

- **Cron-Based Scheduling** – Schedule jobs using standard cron expressions
- **Dynamic Job Registration** – Register jobs via API at runtime
- **At-Least-Once Execution** – Ensures every job runs at least once per schedule
- **Retry Mechanism** – Automatically retries failed jobs
- **Job Expiry & Timeouts** – Jobs expire after a set time window
- **Observability** – View metrics, logs, and execution status
- **Job Dependencies** – Define job execution order
- **Manual Triggers** – Execute jobs instantly via API

---

## 📁 Project Structure

```text
orbit-queue/
├── api/              # HTTP routes and request handling
├── scheduler/        # Core scheduling logic using cron
├── internal/         # Retry logic, job status management, etc.
├── config/           # Environment and application configs
├── cmd/              # Application entrypoints
└── main.go           # Main startup script
```

---

## 🚀 Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/your-username/orbit-queue.git
cd orbit-queue
```

### 2. Run the App

```bash
go run cmd/main.go
```

### 3. Register a Job

```bash
curl -X POST http://localhost:8080/jobs \
  -H "Content-Type: application/json" \
  -d '{
    "id": "job-1",
    "schedule": "*/2 * * * *",
    "command": "echo Hello"
}'
```

---

## 🛠 Tech Stack

- **Language:** Go
- **Web Framework:** Gin
- **Scheduler:** robfig/cron

**Planned:**

- PostgreSQL (for persistence)
- Redis (for distributed queueing)
- Prometheus & Grafana (for metrics)
- Web dashboard (React or Svelte)

---

## 📈 Roadmap

| Feature                        | Branch Name                |
|---------------------------------|---------------------------|
| Initial project setup           | `feature/initial-setup`   |
| HTTP server + health check      | `feature/api-server`      |
| Core scheduler setup            | `feature/scheduler-core`  |
| Cron-based job execution        | `feature/cron-jobs`       |
| Dynamic job registration API    | `feature/dynamic-jobs-api`|
| Job persistence (DB)            | `feature/job-db`          |
| Manual trigger support          | `feature/manual-trigger`  |
| Retry mechanism                 | `feature/job-retry`       |
| Job dependencies/DAG            | `feature/job-dependencies`|
| Observability & metrics (Prom)  | `feature/observability`   |
| Log monitoring per job          | `feature/log-monitoring`  |
| Dashboard UI                    | `feature/dashboard-ui`    |

---

## 📄 License

MIT License

---
