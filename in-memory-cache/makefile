dependencies:
	go mod tidy
	go mod vendor

benchmark:
	go test -bench=. -benchmem -count 5 -benchtime=100000x > results/results.out

generate-graph:
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -OutputFileTemplate=get_%s -LegendTemplate="{1}" -OptionsTemplate={0} -KeyTemplate={0}/{1}/{2} -Filter=functionName!=\"Get\" -TableWidth=2048 -TableLegendSpan=2
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -OutputFileTemplate=set_%s -LegendTemplate="{1}" -OptionsTemplate={0} -KeyTemplate={0}/{1}/{2} -Filter=functionName!=\"Set\" -TableWidth=2048 -TableLegendSpan=2
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -OutputFileTemplate=remove_%s -LegendTemplate="{1}" -OptionsTemplate={0} -KeyTemplate={0}/{1}/{2} -Filter=functionName!=\"Remove\" -TableWidth=2048 -TableLegendSpan=2
