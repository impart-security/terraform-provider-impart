default: install

# Acceptance test configs use source = "impart-security/impart", so the provider under
# test must be registered under that namespace on both Terraform and OpenTofu.
TF_ACC_PROVIDER_NAMESPACE ?= impart-security
export TF_ACC_PROVIDER_NAMESPACE

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

# Same suite driven by the OpenTofu CLI. Short provider sources resolve to
# registry.opentofu.org under tofu, so the provider host must match.
testacc-tofu:
	$(info ************  Running Acceptance Tests (OpenTofu)  ************)
	@# An empty TF_ACC_TERRAFORM_PATH would silently fall back to terraform.
	@command -v tofu >/dev/null || { echo "error: tofu not found on PATH" >&2; exit 2; }
	TF_ACC=1 TF_ACC_TERRAFORM_PATH="$$(command -v tofu)" TF_ACC_PROVIDER_HOST=registry.opentofu.org \
		go test $(TESTARGS) -count=1 -parallel=4 -timeout 120m -v ./...

# End-to-end check of the published release under OpenTofu (needs IMPART_TOKEN; creates
# and destroys one label). MIGRATE=1 applies with terraform and takes over with tofu.
smoke-tofu:
	$(info ************  OpenTofu smoke test  ************)
	./scripts/opentofu-smoke.sh

validate-examples:
	$(info ************  Validating Examples ($(or $(CLI),terraform))  ************)
	CLI=$(or $(CLI),terraform) ./scripts/validate-examples.sh

package:
	$(info ************ Build and package  ************)
	GOOS=$(os) GOARCH=$(arch) go build -buildvcs=false -trimpath -ldflags="-X 'main.version=v$(version)'" -o terraform-provider-impart_v$(version) . && \
	zip terraform-provider-impart_$(version)_$(os)_$(arch).zip terraform-provider-impart_v$(version)
