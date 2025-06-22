build_image:
	docker build -t cevataykans/nwdaf:latest --target release .
	docker login
	docker push cevataykans/nwdaf:latest