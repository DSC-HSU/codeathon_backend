colon := :
$(colon) := :

prepare:
	@echo "Please install the following dependencies:"
	@echo "1. Install supabase cli:"
	@echo "        https://supabase.com/docs/guides/cli/getting-started"
	@echo "2. Install weaver:"
	@echo "        go install github.com/ServiceWeaver/weaver/cmd/weaver@latest"
	@echo "3. Install swaggo:"
	@echo "        https://github.com/swaggo/gin-swagger"
	@echo "If you have installed all the dependencies, just wait for the supabase spin up and run dev-start"
	supabase start

start-rabbitmq:
	docker run -it --rm --name rabbitmq -p 5672$(:)5672 -p 15672$(:)15672 rabbitmq$(:)4.0-management

dev-start:
	weaver generate ./...
	go run ./cmd/server

save-dev-data:
	supabase db pull --schema auth,storage --local

load-dev-data:
	supabase migration up --local

dev-stop:
	supabase stop

dev-start-config:
	weaver generate ./...
	SERVICEWEAVER_CONFIG=weaver.toml go run ./cmd/server