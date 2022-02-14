# vars
DC=docker-compose
DCTEST=$(DC) -f docker-compose.yml -f docker-compose.test.yml
# DCTEST=$(DC) -f docker-compose.test.yml

build-api-docker-image:
	$(DC) build api

# run continuous integration tests
# remember to add steps for removing volumes if they get added
ci-test:
	$(DCTEST) down --remove-orphans
	$(DCTEST) up \
		--build \
		--exit-code-from ci-test \
		ci-test
	$(DCTEST) down
