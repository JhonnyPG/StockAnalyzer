# Stock Analyzer v1.0.0

A full-stack application for analyzing stock recommendations and market insights. The system fetches stock data from an external API, processes it, and provides actionable recommendations based on various metrics.

## 🚀 Features

- Real-time stock data analysis
- Broker recommendation tracking
- Price target monitoring
- Rating change analysis
- Best stock recommendations
  
## 🏗 Architecture

The project follows a microservices architecture with two main components:

### Backend (Go)

Located in `/backend` directory:
- RESTful API service
- Database operations
- External API integration
- Business logic processing

### Frontend (Vue.js)

Located in `/frontend/stock-ui` directory:
- Interactive UI
- Real-time data display
- Responsive design
- State management with Pinia

## 📋 Requirements

### Backend Requirements
- Go 1.22.2 or higher
- CockroachDB (Cloud hosted)
- Access to the stock recommendations API

### Frontend Requirements
- Node.js 18.x or higher
- npm 9.x or higher
- Modern web browser with JavaScript enabled

## 🛠 Dependencies

### Backend Dependencies
```go
require github.com/lib/pq v1.10.9 // PostgreSQL driver for CockroachDB
```

### Frontend Dependencies
```json
{
  "dependencies": {
    "@heroicons/vue": "^2.2.0",
    "@tailwindcss/vite": "^4.0.15",
    "axios": "^1.8.4",
    "pinia": "^3.0.1",
    "tailwindcss": "^4.0.15",
    "vue": "^3.5.13",
    "vue-router": "^4.5.0"
  },
  "devDependencies": {
    "@tsconfig/node22": "^22.0.0",
    "@types/node": "^22.13.9",
    "@vitejs/plugin-vue": "^5.2.1",
    "@vue/eslint-config-prettier": "^10.2.0",
    "@vue/eslint-config-typescript": "^14.5.0",
    "@vue/tsconfig": "^0.7.0",
    "eslint": "^9.21.0",
    "eslint-plugin-vue": "~10.0.0",
    "jiti": "^2.4.2",
    "npm-run-all2": "^7.0.2",
    "prettier": "3.5.3",
    "typescript": "~5.8.0",
    "vite": "^6.2.1",
    "vite-plugin-vue-devtools": "^7.7.2",
    "vue-tsc": "^2.2.8"
  }
}
```

## 🚦 API Endpoints

### GET /stocks
Returns a list of all stocks with their latest recommendations.

**Response:**
```json
[
  {
    "ticker": "AAPL",
    "company": "Apple Inc",
    "brokerage": "Morgan Stanley",
    "action": "Upgrade",
    "rating_from": "Hold",
    "rating_to": "Buy",
    "target_from": "$150",
    "target_to": "$180",
    "time": "2024-03-20T10:30:00Z"
  }
]
```

### GET /recommendations
Returns the best stock recommendation based on multiple factors.

**Response:**
```json
{
  "ticker": "TSLA",
  "company": "Tesla Inc",
  "brokerage": "JP Morgan",
  "action": "Upgrade",
  "rating_from": "Hold",
  "rating_to": "Strong Buy",
  "target_from": "$200",
  "target_to": "$300",
  "time": "2024-03-20T15:45:00Z"
}
```

## 🧪 Testing

### Backend Tests
The backend includes comprehensive unit tests using mocks:

- Database operation tests (`/backend/internal/database/`)
  - Stock insertion tests
  - Error handling tests

Run backend tests:
```bash
cd backend
go test ./internal/... -v
```


## 🚀 Getting Started

1. Clone the repository
2. Set up the backend:
- First run "command/server/main.go" to insert the api data into the DB.
- then run the "command/api/main.go" to start the local api  in listening mode.
```bash
cd backend
go mod download
go run command/server/main.go
go run command/api/main.go
```

3. Set up the frontend:
```bash
cd frontend/stock-ui
npm install
npm run dev
```

4. Access the application at `http://localhost:5173`

## 📦 Project Structure

```
.
├── backend/
│   ├── command/
│   │   ├── api/        # API server entry point
│   │   └── server/     # Data fetching service
│   ├── internal/
│   │   ├── api/        # External API client
│   │   ├── database/   # Database operations
│   │   ├── handlers/   # HTTP handlers
│   │   └── models/     # Data models
│   └── go.mod          # Go dependencies
├── frontend/
│   └── stock-ui/
│       ├── src/
│       │   ├── components/  # Vue components
│       │   ├── stores/      # Pinia stores
│       │   ├── router/      # Vue Router config
│       │   └── assets/      # Static assets
│       └── package.json     # Node.js dependencies
└── README.md
```

## 🔄 Version Control

The project follows Semantic Versioning (SemVer):

- **Major version (1.x.x)**: Incompatible API changes
- **Minor version (x.1.x)**: New features in a backward-compatible manner
- **Patch version (x.x.1)**: Backward-compatible bug fixes

Current version: 1.0.0


## 🤝 Thanks.

I learned a lot doing this project, so I am truly grateful for the opportunity to enhance my skills and knowledge. This experience has been invaluable.
