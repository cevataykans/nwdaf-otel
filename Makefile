build_image:
	docker build -t cevataykans/nwdaf:latest --target release .
	docker push cevataykans/nwdaf:latest

generate_external_scaler:
	protoc --go_out=generated/externalscaler --go-grpc_out=generated/externalscaler server/externalscaler/external_scaler.proto

generate_openapi:
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

install-nuc2:
	bash scripts/infra/nuc2/install.sh

uninstall-nuc2:
	bash scripts/infra/nuc2/uninstall.sh

install-nuc1:
	bash scripts/infra/nuc1/install.sh

uninstall-nuc1:
	bash scripts/infra/nuc1/uninstall.sh

start-analytics:
	helm install -f helm/analyticsinfo_values.yaml nwdaf-analytics-info ./helm -n aether-5gc

stop-analytics:
	helm uninstall nwdaf-analytics-info -n aether-5gc

get-plots:
	scp -r sevinc@131.159.25.36:/home/sevinc/graphs ./graphs