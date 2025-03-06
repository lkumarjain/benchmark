dependencies:
	go mod tidy
	go mod vendor

benchmark-producer:
	go test -timeout=5h -bench=Producer -benchmem -count 5 -benchtime=10000x > results/producer.out

benchmark-consumer:
	go test -timeout=5h -bench=Consumer -benchmem -count 5 -benchtime=10000x > results/consumer.out

generate-graph:
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -ParserFileName=producer.out -OutputFileTemplate=AsyncProducer_%s -LegendTemplate="{1}" -OptionsTemplate={2} -KeyTemplate={0}/{1}/{2} -Filter=functionName!~\"@ProduceAsync\" -TableWidth=2048 -TableLegendSpan=2
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -ParserFileName=producer.out -OutputFileTemplate=SyncProducer_%s -LegendTemplate="{1}" -OptionsTemplate={2} -KeyTemplate={0}/{1}/{2} -Filter=functionName!~\"@ProduceSync\" -TableWidth=2048 -TableLegendSpan=2
	${PWD}/../result/bin/result -ParserFilePath=${PWD}/results -ParserFileName=consumer.out -OutputFileTemplate=Consumer_%s -LegendTemplate="{1}" -OptionsTemplate={2} -KeyTemplate={0}/{1}/{2} -TableWidth=2048 -TableLegendSpan=2 
