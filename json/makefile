dependencies:
	go mod tidy
	go mod vendor

benchmark:
	go test -bench=. -benchmem -timeout=5h -count 5 -benchtime=100000x > results/results.out

generate-graph:
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -OutputFileTemplate=Iterate_%s -LegendTemplate="{1}" -OptionsTemplate={0} -KeyTemplate={0}/{1} -Filter=functionName!~\"Iterate*\" -TableWidth=2048 -TableLegendSpan=2
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -OutputFileTemplate=Struct_%s -LegendTemplate="{1}" -OptionsTemplate={0} -KeyTemplate={0}/{1} -Filter=functionName!~\"Struct*\" -TableWidth=2048 -TableLegendSpan=2 
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -OutputFileTemplate=Interface_%s -LegendTemplate="{1}" -OptionsTemplate={0} -KeyTemplate={0}/{1} -Filter=functionName!~\"Interface*\" -TableWidth=2048 -TableLegendSpan=2
