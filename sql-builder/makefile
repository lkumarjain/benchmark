dependencies:
	go mod tidy
	go mod vendor

benchmark:
	go test -timeout=5h -bench=. -benchmem -count 5 -benchtime=1000000x > results/results.out

generate-graph:
	cd ${PWD}/../result/bin && ./result -ParserFilePath=${PWD}/results -OutputFileTemplate=Insert_%s -LegendTemplate={2} -KeyTemplate={0}/{1}/{2} -Filter=scenarioName!=\"Insert\" -TableWidth=2048 -TableLegendSpan=1
	cd ${PWD}/../result/bin && ./result -ParserFilePath=${PWD}/results -OutputFileTemplate=Select_%s -LegendTemplate={2} -KeyTemplate={0}/{1}/{2} -Filter=scenarioName!=\"Select\" -TableWidth=2048 -TableLegendSpan=1
	cd ${PWD}/../result/bin && ./result -ParserFilePath=${PWD}/results -OutputFileTemplate=Update_%s -LegendTemplate={2} -KeyTemplate={0}/{1}/{2} -Filter=scenarioName!=\"Update\" -TableWidth=2048 -TableLegendSpan=1
	cd ${PWD}/../result/bin && ./result -ParserFilePath=${PWD}/results -OutputFileTemplate=Delete_%s -LegendTemplate={2} -KeyTemplate={0}/{1}/{2} -Filter=scenarioName!=\"Delete\" -TableWidth=2048 -TableLegendSpan=1
