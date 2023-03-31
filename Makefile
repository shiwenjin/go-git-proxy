serverName="yc3"
finalName="go-git-proxy"

.PHONY: build

build:
	@docker build -t ${finalName} .
	@docker save -o ${finalName}.tar ${finalName}
	@scp -r ${finalName}.tar  yc3:/go/docker_file/${finalName}.tar
	@ssh ${serverName} "docker load < /go/docker_file/${finalName}.tar"
	@ssh ${serverName} "docker rm -f ${finalName}"
	@ssh ${serverName} "docker run -d --restart=always -p 80:80 --name ${finalName} ${finalName}"
