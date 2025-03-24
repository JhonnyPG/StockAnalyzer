package database

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/JhonnyPG/StockAnalyzer/internal/models"
)

// MockDB implementa una interfaz para mockear las operaciones de base de datos
type MockDB struct {
	execFunc    func(query string, args ...interface{}) (sql.Result, error)
	shouldError bool
	lastQuery   string
	lastArgs    []interface{}
}

// Implementación del método Exec para el mock
func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.lastQuery = query
	m.lastArgs = args

	if m.shouldError {
		return nil, errors.New("mock database error")
	}

	return sql.Result(&mockResult{}), nil
}

// mockResult implementa sql.Result para testing
type mockResult struct{}

func (m *mockResult) LastInsertId() (int64, error) { return 0, nil }
func (m *mockResult) RowsAffected() (int64, error) { return 1, nil }

// TestInsertStock prueba la inserción de stocks con diferentes escenarios
func TestInsertStock(t *testing.T) {
	testTime := time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name        string
		stock       models.Stock
		mockError   bool
		wantErr     bool
		expectedErr string
	}{
		{
			name: "successful insert",
			stock: models.Stock{
				Ticker:     "AAPL",
				Company:    "Apple Inc.",
				Brokerage:  "BrokerX",
				Action:     "Buy",
				RatingFrom: "Hold",
				RatingTo:   "Buy",
				TargetFrom: "150.0",
				TargetTo:   "200.0",
				Time:       testTime,
			},
			mockError: false,
			wantErr:   false,
		},
		{
			name: "database error",
			stock: models.Stock{
				Ticker: "GOOGL",
			},
			mockError:   true,
			wantErr:     true,
			expectedErr: "failed to insert stock: mock database error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &MockDB{shouldError: tt.mockError}

			err := InsertStock(mockDB, tt.stock)

			if (err != nil) != tt.wantErr {
				t.Errorf("InsertStock() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && err.Error() != tt.expectedErr {
				t.Errorf("Expected error message: %s, got: %s", tt.expectedErr, err.Error())
			}

			// Verifica que se llamó al método Exec
			if mockDB.lastQuery == "" {
				t.Error("Exec method was not called")
			}

			// Verifica los argumentos pasados
			expectedArgs := []interface{}{
				tt.stock.Ticker,
				tt.stock.Company,
				tt.stock.Brokerage,
				tt.stock.Action,
				tt.stock.RatingFrom,
				tt.stock.RatingTo,
				tt.stock.TargetFrom,
				tt.stock.TargetTo,
				tt.stock.Time,
			}

			if len(mockDB.lastArgs) != len(expectedArgs) {
				t.Errorf("Expected %d args, got %d", len(expectedArgs), len(mockDB.lastArgs))
			}
		})
	}
}
