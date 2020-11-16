
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## applications

run-server-v1 :		## Run server-v1
	cd server-v1 && make run

run-server-v2 :		## Run server-v2
	cd server-v2 && make run

run-client-v1-server-v1 :		## Run client-v1 connecting to server-v1
	cd client-v1 && make run-to-server-v1

run-client-v1-server-v2 :		## Run client-v1 connecting to server-v2
	cd client-v1 && make run-to-server-v2

run-client-v2-server-v1 :		## Run client-v2 connecting to server-v1
	cd client-v2 && make run-to-server-v1

run-client-v2-server-v2 :		## Run client-v2 connecting to server-v2
	cd client-v2 && make run-to-server-v2

## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo ""
