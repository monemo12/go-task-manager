# Task Manager

一個基於 Go 語言開發的任務管理系統，支持多種 API 接入方式。

## 目錄結構

```
.
├── README.md
├── go.mod
├── go.sum
└── internal/
    ├── domain/             # 領域模型和核心接口定義
    │   └── task.go         # 任務相關的領域模型和接口
    │
    ├── repository/         # 數據存儲層實現
    │   └── memory/         # 內存存儲實現
    │       └── task.go
    │
    ├── service/           # 業務邏輯層
    │   └── task/          # 任務服務實現
    │       └── service.go
    │
    └── delivery/          # API 傳輸層
        ├── rest/          # RESTful API 實現
        │   └── handler.go
        │
        ├── graphql/       # GraphQL API 實現
        │   └── handler.go
        │
        └── grpc/         # gRPC API 實現
            └── handler.go

└── pkg/                  # 可重用的包
    └── validator/        # 驗證器實現
        ├── task.go       # 任務相關的驗證規則
        └── validator.go  # 通用驗證邏輯

```

## 架構說明

- **Domain Layer**: 定義核心業務模型和接口
- **Repository Layer**: 實現數據持久化
- **Service Layer**: 實現核心業務邏輯
- **Delivery Layer**: 提供多種 API 訪問方式
  - RESTful API
  - GraphQL API
  - gRPC API

## 設計原則

本項目遵循 SOLID 原則：

1. **單一職責原則 (SRP)**
   - 每個包都有明確的、單一的職責
   - 各層之間職責劃分清晰

2. **開放封閉原則 (OCP)**
   - 通過接口定義實現擴展
   - 新增功能無需修改現有代碼

3. **里氏替換原則 (LSP)**
   - 所有實現都可以替換其接口
   - 確保系統的可測試性

4. **接口隔離原則 (ISP)**
   - 接口精簡且專注
   - 客戶端只依賴其需要的接口

5. **依賴倒置原則 (DIP)**
   - 高層模塊不依賴低層模塊
   - 都依賴於抽象接口

## 開發計劃

- [x] 基礎架構搭建
- [ ] 核心領域模型實現
- [ ] 內存存儲實現
- [ ] RESTful API 實現
- [ ] GraphQL API 實現
- [ ] gRPC API 實現
- [ ] 單元測試覆蓋
- [ ] 性能測試和優化 