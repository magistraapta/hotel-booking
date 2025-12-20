# Docker Setup with Local Database

This guide explains how to run the backend in a Docker container while connecting to a PostgreSQL database running on your local machine.

## Prerequisites

1. Docker Desktop installed and running
2. PostgreSQL running locally on your laptop
3. Your local database is accessible (not blocked by firewall)

## Setup Steps

### 1. Configure Your Local Database

Make sure your local PostgreSQL is running and accessible. By default, PostgreSQL listens on `localhost:5432`.

### 2. Update Your `.env` File

Create a `.env` file in the `backend/` directory (or copy from `.env.example`):

```bash
# For Docker: use host.docker.internal to connect to your local database
DB_URL="postgres://your_username:your_password@host.docker.internal:5432/your_database_name"

JWT_SECRET_KEY=your_secret_key_here
SERVER_PORT=8080
```

**Important:** 
- Replace `your_username`, `your_password`, and `your_database_name` with your actual database credentials
- Use `host.docker.internal` as the host (this is a special DNS name that Docker provides to access the host machine)
- Use the port where your local PostgreSQL is running (default is `5432`)

### 3. Build and Run with Docker Compose

```bash
cd backend
docker compose up --build
```

Or to run in detached mode:

```bash
docker compose up -d --build
```

### 4. Verify Connection

The backend should now be running in Docker and connected to your local database. You can:

- Check logs: `docker compose logs -f server`
- Test the API: `curl http://localhost:8080/health` (or your health endpoint)

## Troubleshooting

### Connection Refused

If you get connection errors:

1. **Check PostgreSQL is running:**
   ```bash
   # macOS
   brew services list
   # or
   pg_isready
   ```

2. **Verify PostgreSQL accepts connections:**
   - Check `postgresql.conf` - ensure `listen_addresses = '*'` or includes your network
   - Check `pg_hba.conf` - ensure it allows connections from Docker network

3. **For Linux users:**
   - `host.docker.internal` might not work. Use your host machine's IP address instead:
     ```bash
     # Find your host IP
     ip addr show docker0 | grep inet
     # Or use the gateway IP
     docker network inspect bridge | grep Gateway
     ```

### Alternative: Use Host Network Mode (Linux only)

On Linux, you can use host network mode to directly access localhost:

```yaml
# In compose.yaml, add:
network_mode: "host"
```

Then use `localhost` or `127.0.0.1` in your DB_URL instead of `host.docker.internal`.

## Building Docker Image Only

If you just want to build the image without running it:

```bash
docker build -t backend-app .
```

Then run it manually:

```bash
docker run -p 8080:8080 \
  -e DB_URL="postgres://user:pass@host.docker.internal:5432/dbname" \
  -e JWT_SECRET_KEY="your_secret" \
  -e SERVER_PORT=8080 \
  --add-host=host.docker.internal:host-gateway \
  backend-app
```

