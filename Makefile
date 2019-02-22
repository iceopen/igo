default: release

release:
	@rm -f bin/*
	@/bin/bash ./scripts/release.sh

clean:
	@echo "--> Cleaning..."
	@rm -f bin/*