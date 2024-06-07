package auction_test

import (
	"context"
	"fmt"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database/auction"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCloseAuctionAutomatically(t *testing.T) {
	os.Setenv("MONGO_INITDB_ROOT_USERNAME", "admin")
	os.Setenv("MONGO_INITDB_ROOT_PASSWORD", "admin")
	os.Setenv("MONGODB_URL", "mongodb://admin:admin@172.31.0.11:27017/auctions?authSource=admin")
	os.Setenv("MONGODB_DB", "auctions")
	os.Setenv("AUCTION_DURATION", "60")

	fmt.Println("MONGODB_URL:", os.Getenv("MONGODB_URL"))
	fmt.Println("MONGODB_DB:", os.Getenv("MONGODB_DB"))

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error trying to connect the mongodb", err)
	}
	assert.NoError(t, err)
	defer func() {
		err := client.Disconnect(context.Background())
		assert.NoError(t, err)
	}()

	// Verificar a conexão com o MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Error trying to ping the mongodb", err)
	}
	assert.NoError(t, err)

	// Acessa o banco de dados de testes
	db := client.Database(os.Getenv("MONGODB_DB"))

	// Cria o repositório de leilões
	repo := auction.NewAuctionRepository(db)

	// Define os parâmetros do leilão
	auctionId := uuid.New().String()
	auctionEntity := &auction_entity.Auction{
		Id:          auctionId,
		ProductName: "Test Product",
		Category:    "Test Category",
		Description: "Test Description",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
		Duration:    60 * time.Second, // Define a duração do leilão para 60 segundos
	}

	// Tentar criar o leilão no banco de dados
	if err := repo.CreateAuction(context.Background(), auctionEntity); err != nil {
		fmt.Println("CreateAuction Error")
		fmt.Println(err)
	}
	assert.NoError(t, err)

	// Aguarda mais do que a duração do leilão para validar o objetivo do desafio que é o fechamento automático
	time.Sleep(65 * time.Second)
	var result auction.AuctionEntityMongo
	err = repo.Collection.FindOne(context.Background(), bson.M{"_id": auctionId}).Decode(&result)
	assert.NoError(t, err)

	assert.Equal(t, auction_entity.Completed, result.Status)
}
