package database

import "github/nothiaki/idempotency/internal/models"

var Data = []models.Order{
  { Id: 0, IdempotencyId: "598cd438-9044-403f-85f0-9cb5cd9c3a25", Name: "pizza", Price: 10.90 },
  { Id: 1, IdempotencyId: "4c98f7a1-68f2-4ed8-b7ed-10141c1c900b", Name: "hamburger", Price: 4.90 },
  { Id: 2, IdempotencyId: "ba2ace54-4dc4-4e9c-b285-6282a1e62366", Name: "cola", Price: 1.50 },
  { Id: 3, IdempotencyId: "17cb50a0-9428-483b-aab7-b87cef3bad4a", Name: "juice", Price: 1.00 },
}
