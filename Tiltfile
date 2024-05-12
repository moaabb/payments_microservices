load('ext://restart_process', 'docker_build_with_restart')

local_resource(
  'customer_svc_app',
  'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/customer_svc ./',
  deps=['./main.go', './routes.go'])

docker_compose('docker-compose.yaml')
  
docker_build(
  'customer_svc',
  '.',
  dockerfile='./Dockerfile',
  only=[
    './output',
  ],
  live_update=[
    sync('./output', '/app/')
  ],
  )