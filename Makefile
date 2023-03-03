build-backend:
		@echo "Building backend..."
		@cd backend && go build -o ../bin/puppet-installer
build-frontend:
		@echo "Building frontend..."
		@cd ui && npm run build

build: build-backend build-frontend

build-images:
		@echo "Building images..."
		@docker build -t backend:latest -f backend/Dockerfile .
		@docker build -t frontend:latest -f frontend/Dockerfile .

run:
		@echo "Running..."
		@docker-compose up