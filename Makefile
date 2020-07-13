.PHONY: %

cloud-sql-proxy:
	docker run --rm -it \
		-v ${HOME}/.config/gcloud/application_default_credentials.json:/credential.json \
		-p 127.0.0.1:3306:3306 \
		gcr.io/cloudsql-docker/gce-proxy:1.17 /cloud_sql_proxy \
		-instances=${CLOUD_SQL_CONNECTION_NAME}=tcp:0.0.0.0:3306 -credential_file=/credential.json
