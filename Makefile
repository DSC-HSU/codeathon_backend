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
dev-start:
	weaver generate ./...
	go run ./cmd/server

save-dev-data:
	supabase db pull --schema auth,storage --local

load-dev-data:
	supabase migration up

dev-stop:
	supabase stop