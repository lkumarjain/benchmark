dependencies:
	go mod tidy
	go mod vendor

benchmark:
	go test -bench=. -benchmem -count 5 -benchtime=10000x > results/results.out

generate-graph:
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -OutputFileTemplate=compile_%s -LegendTemplate="{1} {2}" -KeyTemplate={0}/{1}/{2} -Filter=functionName==\"Evaluate\" -TableWidth=2048 -TableLegendSpan=2
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -OutputFileTemplate=evaluate_%s -LegendTemplate="{1} {2}" -KeyTemplate={0}/{1}/{2} -Filter=functionName==\"Compile\" -TableWidth=2048 -TableLegendSpan=2
