default: install

generate:
	$(info ************  Generating Provider Documentation  ************)
	go generate ./...

install:
	$(info ************  Building and Installing Binary  ************)
	go install .

test:
	$(info ************  Running Unit Tests  ************)
	go test -v -count=1 -parallel=4 ./...

testacc:
	$(info ************  Running Acceptance Tests  ************)
	TF_ACC=1 go test $(TESTARGS) -count=1 -parallel=4 -timeout 120m -v ./...

package:
	$(info ************ Build and package  ************)
	GOOS=$(os) GOARCH=$(arch) go build -buildvcs=false -trimpath -ldflags="-X 'main.version=v$(version)'" -o terraform-provider-impart_v$(version) . && \
	zip terraform-provider-impart_$(version)_$(os)_$(arch).zip terraform-provider-impart_v$(version)
