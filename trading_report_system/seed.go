package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	db := client.Database("trading_db")

	// 1. Seed Strategy Info
	siColl := db.Collection(CollStrategyInfo)
	siColl.Drop(ctx)
	si := StrategyInfo{
		ID:           "abcd1122",
		StrategyName: "Straddle",
		Category:     "option_buying",
		RiskLevel:    "high",
	}
	siColl.InsertOne(ctx, si)

	// 2. Seed Positions
	posColl := db.Collection(CollPositions)
	posColl.Drop(ctx)
	pos := Position{
		ClientID:     "C101",
		StrategyName: "Straddle",
		Symbol:       "NIFTY24NOV18000CE",
		NetQuantity:  100,
		AvgPrice:     87.9,
	}
	posColl.InsertOne(ctx, pos)

	// 3. Seed Orders and Trades
	orderColl := db.Collection(CollOrders)
	tradeColl := db.Collection(CollTrades)
	orderColl.Drop(ctx)
	tradeColl.Drop(ctx)

	orderID := primitive.NewObjectID()
	order := Order{
		ID:           orderID,
		ClientID:     "C101",
		StrategyName: "Straddle",
		Symbol:       "NIFTY24NOV18000CE",
		OrderType:    "BUY",
		Quantity:     50,
		Price:        88.5,
		CreatedAt:    time.Now(),
	}
	orderColl.InsertOne(ctx, order)

	trade := Trade{
		OrderID:   orderID.Hex(),
		TradeType: "BUY",
		Quantity:  50,
		Price:     88.7,
		Timestamp: time.Now(),
	}
	tradeColl.InsertOne(ctx, trade)

	fmt.Println("Mock data seeded successfully!")
}
