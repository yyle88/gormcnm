COVERAGE_DIR ?= .coverage.out

# cp from: https://github.com/yyle88/osexec/blob/5d35ce11e097573b53d03744479836fb7fdd7e85/Makefile#L4
test:
	@if [ -d $(COVERAGE_DIR) ]; then rm -r $(COVERAGE_DIR); fi
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
