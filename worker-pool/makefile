dependencies:
	go mod tidy
	go mod vendor

benchmark:
	go test -timeout=5h -bench=. -benchmem -count 5 -benchtime=1000000x > results/results.out

generate-graph:
	cd ${PWD}/../result/bin && ./result -ParserFilePath=${PWD}/results -OutputFileTemplate=concurrency_%s -LegendTemplate={2} -KeyTemplate={0}/{2}  -TableWidth=2048 -TableLegendSpan=1
	cd ${PWD}/../result/bin && ./result -ParserFilePath=${PWD}/results -OutputFileTemplate=data_%s -LegendTemplate={1} -KeyTemplate={0}/{1}  -TableWidth=2048 -TableLegendSpan=1