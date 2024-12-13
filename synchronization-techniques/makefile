dependencies:
	go mod tidy
	go mod vendor

benchmark:
	mkdir -p results
	go test -bench=. -benchmem -count 5 -benchtime=100000x > results/results.out

generate-graph:
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -OutputFileTemplate=concurrency_%s -LegendTemplate="{2}" -OptionsTemplate={1} -KeyTemplate={1}/{2} -TableWidth=2348 -TableLegendSpan=1
