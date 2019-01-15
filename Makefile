default: release

release:
	@/bin/bash ./scripts/release.sh

clean:
	@echo "--> Cleaning..."
	@rm -f bin/*