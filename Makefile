LATEST_TAG=v1.0.0-hyo

docker_build_new_minikube:
	# need to change version, such as: 0.0.5
	$(eval $(minikube -p minikube docker-env))
	@docker build -t kubectl-flame-asyncprofiler2:$(LATEST_TAG) .

release_without_publish:
	# need to tag it, such as "git tag -a v0.0.2 -m 'first next release'"
	git tag -a $(LATEST_TAG) -m 'bump tag'
	goreleaser release --rm-dist --skip-publish

release:
	# need to tag it, such as "git tag -a v0.0.2 -m 'first next release'"
	git tag -a $(LATEST_TAG) -m 'bump tag'
	goreleaser release --rm-dist

execute:
	./dist/kubectl-flame_darwin_amd64/kubectl-flame $(TARGET_POD) -n default -t 5m --lang java -f ./output/flamegraph.jfr

expose_hyotestgrails242_service_port:
	kubectl port-forward service/hyotestgrails242 8080:8080

docker_build:
	@docker build -t kubectl-flame-asyncprofiler2:$(LATEST_TAG) .

push_to_docker_hub:
	# run docker_build if not having the image yet
	docker tag kubectl-flame-asyncprofiler2:$(LATEST_TAG) thanhhungle/kubectl-flame-asyncprofiler2:$(LATEST_TAG)
	docker push thanhhungle/kubectl-flame-asyncprofiler2:$(LATEST_TAG)
# README: use from top to bottom, to try it out. can share with SRE the file, to run execute when needed with the pod name input