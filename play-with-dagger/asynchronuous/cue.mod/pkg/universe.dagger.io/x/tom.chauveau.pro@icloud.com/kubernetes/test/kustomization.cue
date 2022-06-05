package kubernetes

import (
	"dagger.io/dagger"

	"universe.dagger.io/x/tom.chauveau.pro@icloud.com/kubernetes"
)

dagger.#Plan & {
	client: filesystem: "./data/hello-kustomize": read: contents: dagger.#FS

	actions: test: {
		_source: client.filesystem."./data/hello-kustomize".read.contents

		url: kustomization: kubernetes.#Kustomize & {
			type: "url"
			url:  "https://github.com/kubernetes-sigs/kustomize.git/examples/helloWorld?ref=v1.0.6"
		}
	}
}
