
 docker_build_new:
 	eval $(minikube -p minikube docker-env)
 	@docker build -t kubectl-flame-asyncprofiler2:v0.0.1 .