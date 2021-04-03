generate:
	openapi-generator generate -i ./test.yaml -g go -o ./test --additional-properties=useOneOfDiscriminatorLookup=true