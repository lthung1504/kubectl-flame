
docker_build_new:
	# need to change version, such as: 0.0.5
	$(eval $(minikube -p minikube docker-env))
	@docker build -t kubectl-flame-asyncprofiler2:v0.0.4 .

release:
	# need to tag it, such as "git tag -a v0.0.2 -m 'first next release'"
	goreleaser release --rm-dist