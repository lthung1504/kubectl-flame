
docker_build_new:
	# need to change version, such as: 0.0.5
	$(eval $(minikube -p minikube docker-env))
	@docker build -t kubectl-flame-asyncprofiler2:v0.0.3 .

release_without_publish:
	# need to tag it, such as "git tag -a v0.0.2 -m 'first next release'"
	git tag -a $(NEW_TAG) -m 'bump tag'
	goreleaser release --rm-dist --skip-publish

release:
	# need to tag it, such as "git tag -a v0.0.2 -m 'first next release'"
	git tag -a $(NEW_TAG) -m 'bump tag'
	goreleaser release --rm-dist

execute:
	./dist/kubectl-flame_darwin_amd64/kubectl-flame $(TARGET_POD) -n default -t 5m --lang java -f ./output/flamegraph.jfr

expose_hyotestgrails242_service_port:
	kubectl port-forward service/hyotestgrails242 8080:8080

push_to_docker_hub:
	# build latest image file, then tag it, then push to hub
	docker tag kubectl-flame-asyncprofiler2:v1.0.0-hyo thanhhungle/kubectl-flame-asyncprofiler2:v1.0.0-hyo
	docker push thanhhungle/kubectl-flame-asyncprofiler2:v1.0.0-hyo
# README: use from top to bottom, to try it out. can share with SRE the file, to run execute when needed with the pod name input