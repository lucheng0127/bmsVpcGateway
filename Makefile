tools:
	go get k8s.io/code-generator/cmd/conversion-gen
	go get k8s.io/code-generator/cmd/deepcopy-gen
	go get k8s.io/code-generator/cmd/defaulter-gen
	go get k8s.io/code-generator/cmd/applyconfiguration-gen/generators@v0.32.0
	go get k8s.io/code-generator/cmd/applyconfiguration-gen/generators@v0.32.0
	go get k8s.io/code-generator/cmd/applyconfiguration-gen/generators@v0.32.0
	go get k8s.io/code-generator/cmd/applyconfiguration-gen/generators@v0.32.0

generate: tools
	./hack/codegen.sh

manifest:
	controller-gen crd paths=./pkg/apis/... output:crd:dir=config/crd