package models

type Order struct{
  Id int8 `json:"id"`
  IdempotencyId string `json:"idempotency_id"`
  Name string `json:"name"`
  Price float32 `json:"price"`
}
