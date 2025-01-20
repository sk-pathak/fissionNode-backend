-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (name, email, username, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1 RETURNING *;


-- name: GetNodes :many
SELECT * FROM nodes;

-- name: GetNodeByID :one
SELECT * FROM nodes WHERE id = $1;

-- name: RegisterNode :one
INSERT INTO nodes (node_name, ip_address, capacity, status) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: DeleteNode :exec
DELETE FROM nodes WHERE id = $1 RETURNING *;

-- name: UpdateStatus :exec
UPDATE nodes SET status = $1, last_heartbeat = $2 WHERE id = $3 RETURNING *;

-- Name: GetOnlineNodes :many
SELECT * FROM nodes WHERE status = 'online';

-- Name: GetStaleNodes :many
SELECT * FROM nodes WHERE last_heartbeat < NOW() - INTERVAL '5 minutes';


-- name: GetUserServices :many
SELECT * FROM services WHERE user_id = $1;

-- Name: DeployService :one
INSERT INTO services (user_id, node_id, image, status, public_url) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetServiceByID :one
SELECT * FROM services WHERE id = $1;

-- name: DeleteService :exec
DELETE FROM services WHERE id = $1 RETURNING *;

-- Name: UpdateServiceStatus :exec
UPDATE services SET status = $1 WHERE id = $2;

-- Name: InsertNodeHealthLog :exec
INSERT INTO node_health_logs (node_id, cpu_usage, memory_usage) VALUES ($1, $2, $3);

-- Name: GetNodeHealthLogs :many
SELECT * FROM node_health_logs WHERE node_id = $1 ORDER BY created_at DESC LIMIT $2;


-- Name: GetRunningServicesWithNodes :many
SELECT s.id AS service_id, s.image, s.status, s.public_url, n.node_name, n.ip_address FROM services s JOIN nodes n ON s.node_id = n.id WHERE s.status = 'running';

-- Name: GetNodeUtilization :one
SELECT n.node_name, 
       COUNT(s.id) AS running_services, 
       AVG(nh.cpu_usage) AS avg_cpu_usage, 
       AVG(nh.memory_usage) AS avg_memory_usage 
FROM nodes n 
LEFT JOIN services s ON n.id = s.node_id AND s.status = 'running' 
LEFT JOIN node_health_logs nh ON n.id = nh.node_id 
WHERE n.id = $1 
GROUP BY n.node_name;
