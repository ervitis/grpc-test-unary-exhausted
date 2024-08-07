PHONY: gen

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

generate: ## Generate automated code
	docker run -v `pwd`:/defs namely/protoc-all -f ./proto/*.proto -l go && \
	cp gen/pb-go/pb_impl/* pb_impl && \
	rm -rf gen