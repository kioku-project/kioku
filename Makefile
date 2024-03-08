kind-up:
	./hacks/kind-up.sh

deploy:
	./hacks/deploy.sh

kind-deploy:
	docker compose -f ./docker-compose.build.yml build
	kind load docker-image kioku-frontend:latest kioku-carddeck_service:latest kioku-frontend_proxy:latest kioku-srs_service:latest kioku-user_service:latest kioku-collaboration_service:latest kioku-notification_service:latest
	./hacks/deploy.sh ./hacks/values_local.yaml