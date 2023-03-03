build-backend:
		@echo "Building backend..."
		@cd backend && go build -o ../bin/puppet-installer
build-frontend:
		@echo "Building frontend..."
		@cd ui && npm run build

build: build-frontend build-backend

build-image:
		@echo "Building image..."
		@docker build -t installer:latest -f backend/Dockerfile .
