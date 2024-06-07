build:
	docker-compose up -d --build
	@echo "Aplicativo compilado e iniciado com sucesso."

up:
	docker-compose up -d
	@echo "Aplicativo iniciado com sucesso."

down:
	docker-compose down
	@echo "Contêineres do aplicativo parados e removidos."

tests:
	go test -v -run ^TestCloseAuctionAutomatically$ fullcycle-auction_go/internal/infra/database/auction
	@echo "Testes executados com sucesso."

logs:
	docker-compose logs app

.PHONY: go