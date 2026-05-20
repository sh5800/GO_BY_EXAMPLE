package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Collection Names
const (
	CollOrders       = "orders"
	CollTrades       = "trades"
	CollPositions    = "positions"
	CollStrategyInfo = "strategy_info"
)

// MongoDB Records

type Order struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ClientID     string             `bson:"client_id" json:"client_id"`
	StrategyName string             `bson:"strategy_name" json:"strategy_name"`
	Symbol       string             `bson:"symbol" json:"symbol"`
	OrderType    string             `bson:"order_type" json:"order_type"` // BUY/SELL
	Quantity     int                `bson:"quantity" json:"quantity"`
	Price        float64            `bson:"price" json:"price"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
}

type Trade struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID   string             `bson:"order_id" json:"order_id"`
	TradeType string             `bson:"trade_type" json:"trade_type"` // BUY/SELL
	Quantity  int                `bson:"quantity" json:"quantity"`
	Price     float64            `bson:"price" json:"price"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}

type Position struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ClientID     string             `bson:"client_id" json:"client_id"`
	StrategyName string             `bson:"strategy_name" json:"strategy_name"`
	Symbol       string             `bson:"symbol" json:"symbol"`
	NetQuantity  int                `bson:"net_quantity" json:"net_quantity"`
	AvgPrice     float64            `bson:"avg_price" json:"avg_price"`
}

type StrategyInfo struct {
	ID           string `bson:"_id,omitempty" json:"id"`
	StrategyName string `bson:"strategy_name" json:"strategy_name"`
	Category     string `bson:"category" json:"category"`
	RiskLevel    string `bson:"risk_level" json:"risk_level"`
}

// API Request/Response

type ReportFilters struct {
	ClientIDs        []string `json:"client_ids"`
	StrategyNames    []string `json:"strategy_names"`
	Symbols          []string `json:"symbols"`
	DateFrom         string   `json:"date_from"`
	DateTo           string   `json:"date_to"`
	MinQuantity      *int     `json:"min_quantity"`
	MaxQuantity      *int     `json:"max_quantity"`
	MinPrice         *float64 `json:"min_price"`
	MaxPrice         *float64 `json:"max_price"`
	MinPnL           *float64 `json:"min_pnl"`
	MaxPnL           *float64 `json:"max_pnl"`
	OrderType        string   `json:"order_type"`
	Page             int      `json:"page"`
	Limit            int      `json:"limit"`
	SortField        string   `json:"sort_field"`
	SortDir          int      `json:"sort_dir"` // 1 or -1
	IncludePositions bool     `json:"include_positions"`
	IncludeTrades    bool     `json:"include_trades"`
}

type ReportSummary struct {
	TotalBuy    int     `json:"total_buy"`
	TotalSell   int     `json:"total_sell"`
	NetQuantity int     `json:"net_quantity"`
	Turnover    float64 `json:"turnover"`
	PnL         float64 `json:"pnl"`
}

type DetailedTrade struct {
	TradeID   string    `json:"trade_id"`
	Qty       int       `json:"qty"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

type ClientReport struct {
	ClientID     string          `json:"client_id"`
	StrategyName string          `json:"strategy_name"`
	Symbol       string          `json:"symbol"`
	Summary      ReportSummary   `json:"summary"`
	StrategyInfo *StrategyInfo   `json:"strategy_info,omitempty"`
	Positions    *Position       `json:"positions,omitempty"`
	DetailedData []DetailedTrade `json:"detailed_data"`
}

type PaginatedResponse struct {
	Summary      ReportSummary  `json:"summary"`
	DetailedData []ClientReport `json:"detailed_data"`
}
