# CloudTask - Microservices Task Management Platform

CloudTask is a production-ready, cloud-based microservices task management platform inspired by tools like Trello, Linear, and Notion. It demonstrates modern software engineering and DevOps patterns suitable for an enterprise-level SaaS.

## Architecture & Tech Stack

**Backend:**
- **Go (Golang)**: High-performance microservices (Auth, Task, Team, Notification).
- **PostgreSQL**: Relational database for persistent storage (Users, Teams, Tasks).
- **Redis**: Caching, state management, and Pub/Sub for WebSockets.
- **Microservices Design**: Clean, Layered Architecture (Handler -> Service -> Repository).

**Frontend:**
- **React 18 + TypeScript + Vite**: Fast, modern frontend.
- **Zustand**: Global state management (Auth storage).
- **Vanilla CSS Variables**: Design token system achieving a clean, Notion-like aesthetic without heavy frameworks.

**DevOps & Cloud:**
- **Docker Compose**: Entire stack orchestrated locally.
- **Nginx API Gateway**: Routes client requests (`/api/*`, `/ws/*`) to the respective microservices.
- **GitHub Actions**: Automated CI/CD for testing and building Go and React apps.

---

## Directory Structure

\`\`\`bash
/api-gateway                # Nginx configuration
/infra                      # General Infrastructure docs
/services
  ├── auth-service          # Go Fiber + PostgreSQL + JWT
  ├── task-service          # Go Fiber + PostgreSQL + Core Logic
  ├── team-service          # Go Fiber + PostgreSQL (Team Management)
  └── notification-service  # Go Fiber + WebSockets + Redis Pub/Sub
/frontend                   # Vite React + TS Application
docker-compose.yml          # Local container orchestration
.github/workflows/ci.yml    # CI/CD Pipeline
\`\`\`

## Getting Started Locally

### Prerequisites
- Docker & Docker Compose
- Node.js & npm (if running frontend separately)
- Go 1.21+ (if running backend separately)

### Run via Docker Compose

1. Clone the repository and navigate to the project root.
2. Build and spin up the entire cluster:
   \`\`\`bash
   docker-compose up --build -d
   \`\`\`
   
*This will spin up PostgreSQL, Redis, Nginx API Gateway (Port 80), and the 4 Go microservices.*

3. Spin up the frontend:
   \`\`\`bash
   cd frontend
   npm install
   npm run dev
   \`\`\`
   
*Frontend will be available at `http://localhost:5173`.*

---

## API Endpoints List (via Nginx Gateway on port 80)

### Auth Service
- `POST /api/auth/register` - Create a new user (Body: `email`, `password`, `role`)
- `POST /api/auth/login` - Authenticate and receive JWT (Body: `email`, `password`)

### Team Service (Requires JWT `Bearer {token}`)
- `POST /api/teams/` - Create a Team workspace
- `GET /api/teams/` - List user's teams
- `POST /api/teams/:id/members` - Add a member to the team

### Task Service (Requires JWT `Bearer {token}`)
- `POST /api/tasks/` - Create a Task in a project
- `GET /api/tasks/?team_id=1` - Fetch all tasks for a given team

### Notification Service (Requires JWT `?token={token}`)
- `WS /ws/notifications/?token=` - Connect to the WebSocket Pub/Sub stream stream

---

## Cloud Deployment Guide (AWS or GCP)

1. **Database Strategy**: Provision an AWS RDS (PostgreSQL) or GCP Cloud SQL instance.
2. **Caching Strategy**: Provision an AWS ElastiCache (Redis) or Memorystore instance.
3. **Container Hosting**: Deploy the microservices to AWS ECS (Fargate) or GCP Cloud Run using docker images built by GitHub Actions.
4. **API Gateway**: Use an AWS Application Load Balancer (ALB) or Google Cloud Load Balancer with Nginx Ingress or API Gateway to route traffic.
5. **Frontend**: Deploy the static built `frontend/dist` folder to AWS S3 + CloudFront, or Vercel / Netlify.
6. **Secrets & Config**: Manage environment variables (`JWT_SECRET`, `DB_PASSWORD`) via AWS Secrets Manager or Google Secret Manager.
