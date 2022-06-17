.PHONY: import-repos
TMPDIR := $(shell mktemp -d)

repos-download: ## Download external repo urls from Image Builder
	git clone https://github.com/osbuild/image-builder.git --sparse --depth=1 $(TMPDIR)
	cd $(TMPDIR); git sparse-checkout set distributions/ # reduces amount downloaded
	go run ./cmd/external_repos/main.go download $(TMPDIR)/distributions/

repos-import: ## Import External repo urls from Image Builders into the DB
	go run ./cmd/external_repos/main.go import