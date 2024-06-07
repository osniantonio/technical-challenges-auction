package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
	Duration    int64                           `bson:"duration"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) closeAuctionAutomaticallyAfterTime(auctionEntity *auction_entity.Auction) {
	time.Sleep(auctionEntity.Duration)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": auctionEntity.Id, "status": auction_entity.Active}
	update := bson.M{"$set": bson.M{"status": auction_entity.Completed}}

	_, err := ar.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error("Error trying to close auction", err)
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {

	durationEnv := os.Getenv("AUCTION_DURATION")
	duration, err := strconv.ParseInt(durationEnv, 10, 64)
	if err != nil {
		logger.Error("Error parsing duration of the auction", err)
		return internal_error.NewInternalServerError("Error parsing duration of the auction")
	}
	auctionEntity.Duration = time.Duration(duration) * time.Second
	fmt.Println("CreateAuction - auctionEntity.Duration")
	fmt.Println(auctionEntity.Duration)

	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
		Duration:    auctionEntity.Duration.Nanoseconds(),
	}
	_, err = ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		fmt.Println("CreateAuction Error")
		fmt.Println(err)
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	go ar.closeAuctionAutomaticallyAfterTime(auctionEntity)

	return nil
}
