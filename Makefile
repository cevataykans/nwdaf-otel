build_image:
	docker build -t cevataykans/nwdaf:latest --target release .
	docker push cevataykans/nwdaf:latest

generate_all:
	openapi-generator generate -i templates/5G_APIs/TS29520_Nnwdaf_AnalyticsInfo.yaml -g go-server -o ./generated/temp --additional-properties=packageName=analyticsinfo
	cd ./generated/temp
	go mod tidy
	cd ../..
	mv ./generated/temp/go/* ./generated/analyticsinfo
	openapi-generator generate -i templates/5G_APIs/TS29520_Nnwdaf_DataManagement.yaml -g go-server -o ./generated/temp --additional-properties=packageName=datamanagement
	mv ./generated/temp/go/* ./generated/datamanagement
	openapi-generator generate -i templates/5G_APIs/TS29520_Nnwdaf_EventsSubscription.yaml -g go-server -o ./generated/temp --additional-properties=packageName=eventssubscription
	mv ./generated/temp/go/* ./generated/eventssubscription
	openapi-generator generate -i templates/5G_APIs/TS29520_Nnwdaf_MLModelProvision.yaml -g go-server -o ./generated/temp --additional-properties=packageName=mlmodelprovision
	mv ./generated/temp/go/* ./generated/mlmodelprovision