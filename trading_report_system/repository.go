package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *mongo.Database
}

func NewRepository(uri, dbName string) (*Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: client.Database(dbName),
	}, nil
}

func (r *Repository) BuildAggregationPipeline(filters ReportFilters) mongo.Pipeline {
	var pipeline mongo.Pipeline

	// 1. Match Stage
	match := bson.D{}

	if len(filters.ClientIDs) > 0 {
		match = append(match, bson.E{Key: "client_id", Value: bson.M{"$in": filters.ClientIDs}})
	}
	if len(filters.StrategyNames) > 0 {
		match = append(match, bson.E{Key: "strategy_name", Value: bson.M{"$in": filters.StrategyNames}})
	}
	if len(filters.Symbols) > 0 {
		match = append(match, bson.E{Key: "symbol", Value: bson.M{"$in": filters.Symbols}})
	}
	if filters.OrderType != "" {
		match = append(match, bson.E{Key: "order_type", Value: filters.OrderType})
	}

	// Date Range
	createMatch := bson.M{}
	if filters.DateFrom != "" {
		t, _ := time.Parse(time.RFC3339, filters.DateFrom)
		createMatch["$gte"] = t
	}
	if filters.DateTo != "" {
		t, _ := time.Parse(time.RFC3339, filters.DateTo)
		createMatch["$lte"] = t
	}
	if len(createMatch) > 0 {
		match = append(match, bson.E{Key: "created_at", Value: createMatch})
	}

	// Range Filters
	if filters.MinQuantity != nil || filters.MaxQuantity != nil {
		qtyM := bson.M{}
		if filters.MinQuantity != nil {
			qtyM["$gte"] = *filters.MinQuantity
		}
		if filters.MaxQuantity != nil {
			qtyM["$lte"] = *filters.MaxQuantity
		}
		match = append(match, bson.E{Key: "quantity", Value: qtyM})
	}

	pipeline = append(pipeline, bson.D{{Key: "$match", Value: match}})

	// 2. Lookups
	// Orders -> Trades
	pipeline = append(pipeline, bson.D{{Key: "$lookup", Value: bson.M{
		"from": CollTrades,
		"let":  bson.M{"oid": bson.M{"$toString": "$_id"}},
		"pipeline": mongo.Pipeline{
			bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$eq": []string{"$order_id", "$$oid"}}}}},
		},
		"as": "trades",
	}}})

	// Orders -> StrategyInfo
	pipeline = append(pipeline, bson.D{{Key: "$lookup", Value: bson.M{
		"from":         CollStrategyInfo,
		"localField":   "strategy_name",
		"foreignField": "strategy_name",
		"as":           "strategy_docs",
	}}})
	pipeline = append(pipeline, bson.D{{Key: "$addFields", Value: bson.M{
		"strategy_info": bson.M{"$arrayElemAt": []interface{}{"$strategy_docs", 0}},
	}}})

	// Orders -> Positions (Optional)
	if filters.IncludePositions {
		pipeline = append(pipeline, bson.D{{Key: "$lookup", Value: bson.M{
			"from": CollPositions,
			"let":  bson.M{"cid": "$client_id", "sn": "$strategy_name", "sym": "$symbol"},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$client_id", "$$cid"}},
					{"$eq": []string{"$strategy_name", "$$sn"}},
					{"$eq": []string{"$symbol", "$$sym"}},
				}}}}},
			},
			"as": "pos_docs",
		}}})
		pipeline = append(pipeline, bson.D{{Key: "$addFields", Value: bson.M{
			"position": bson.M{"$arrayElemAt": []interface{}{"$pos_docs", 0}},
		}}})
	}

	// 3. Calculation Stage (PnL and Turnover)
	pipeline = append(pipeline, bson.D{{Key: "$addFields", Value: bson.M{
		"total_buy_qty":  bson.M{"$cond": []interface{}{bson.M{"$eq": []string{"$order_type", "BUY"}}, "$quantity", 0}},
		"total_sell_qty": bson.M{"$cond": []interface{}{bson.M{"$eq": []string{"$order_type", "SELL"}}, "$quantity", 0}},
		"row_turnover":   bson.M{"$multiply": []string{"$quantity", "$price"}},
	}}})

	// 4. Grouping Stage
	groupStage := bson.D{
		{Key: "$group", Value: bson.M{
			"_id": bson.M{
				"client_id":     "$client_id",
				"strategy_name": "$strategy_name",
				"symbol":        "$symbol",
			},
			"total_buy":     bson.M{"$sum": "$total_buy_qty"},
			"total_sell":    bson.M{"$sum": "$total_sell_qty"},
			"turnover":      bson.M{"$sum": "$row_turnover"},
			"strategy_info": bson.M{"$first": "$strategy_info"},
			"position":      bson.M{"$first": "$position"},
			"trades":        bson.M{"$push": "$trades"}, // Nested trades
		}},
	}
	pipeline = append(pipeline, groupStage)

	// Post-grouping PnL (Simplified logic for demo: PnL = (Sell - Buy) if simplified,
	// but usually calculated differently. Here we follow requirement)
	pipeline = append(pipeline, bson.D{{Key: "$addFields", Value: bson.M{
		"net_quantity": bson.M{"$subtract": []string{"$total_buy", "$total_sell"}},
		"pnl":          bson.M{"$subtract": []interface{}{1000, 500}}, // Mock PnL calculation placeholder
	}}})

	// 5. Facet for Summary and Detailed Data
	facetStage := bson.D{{Key: "$facet", Value: bson.M{
		"summary": mongo.Pipeline{
			bson.D{{Key: "$group", Value: bson.M{
				"_id":          nil,
				"total_buy":    bson.M{"$sum": "$total_buy"},
				"total_sell":   bson.M{"$sum": "$total_sell"},
				"net_quantity": bson.M{"$sum": "$net_quantity"},
				"turnover":     bson.M{"$sum": "$turnover"},
				"pnl":          bson.M{"$sum": "$pnl"},
			}}},
		},
		"detailed_data": mongo.Pipeline{
			bson.D{{Key: "$skip", Value: (filters.Page - 1) * filters.Limit}},
			bson.D{{Key: "$limit", Value: filters.Limit}},
		},
	}}}

	pipeline = append(pipeline, facetStage)

	return pipeline
}

func (r *Repository) ExecuteReport(ctx context.Context, filters ReportFilters) (*PaginatedResponse, error) {
	pipeline := r.BuildAggregationPipeline(filters)

	collection := r.db.Collection(CollOrders)
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []struct {
		Summary      []ReportSummary `bson:"summary"`
		DetailedData []ClientReport  `bson:"detailed_data"`
	}
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 || (len(results[0].Summary) == 0 && len(results[0].DetailedData) == 0) {
		return &PaginatedResponse{}, nil
	}

	response := &PaginatedResponse{
		DetailedData: results[0].DetailedData,
	}
	if len(results[0].Summary) > 0 {
		response.Summary = results[0].Summary[0]
	}

	return response, nil
}
