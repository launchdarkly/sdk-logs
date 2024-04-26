.PHONY: logdrkly logdrkly-fmt gen-code-schema

gen-code-schema:
	go-jsonschema -p logs definition/schema/ld_log_codes.json > tools/logdrkly/internal/logs/ld_log_codes.go
	# This is to work around an inadequacy in the generated types in how it handles patternProperties.
	sed -i '/type LdLogCodesJsonConditions map\[string\]interface{}/ c\type LdLogCodesJsonConditions map\[string\]Condition' tools/logdrkly/internal/logs/ld_log_codes.go

logdrkly: gen-code-schema
	cd tools/logdrkly && go build .

logdrkly-fmt:
	cd tools/logdrkly && go fmt .

all: logdrkly
