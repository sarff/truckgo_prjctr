package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alexandear/truckgo/order/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to initialize gorm: %v", err)
	}

	return gormDB, mock
}

func TestCreateOrder(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewOrderRepository(db)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "orders" ("number","status","price","user_id","driver_id","is_archived") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), 0, 100.5, 1, 0, false).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	orderID, err := repo.Create(100.5, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if orderID != 1 {
		t.Fatalf("expected order ID 1, got %d", orderID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestUpdateOrder(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewOrderRepository(db)

	order := models.Order{ID: 1, Price: 200.75, UserID: 2}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "orders" SET "price"=$1 WHERE id = $2 AND "id" = $3`)).
		WithArgs(300.99, order.ID, order.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	updates := map[string]any{"price": 300.99}
	err := repo.Update(order, updates)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestFindAllOrders(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewOrderRepository(db)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "orders"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	rows := sqlmock.NewRows([]string{"id", "price", "user_id"}).
		AddRow(1, 50.0, 1).
		AddRow(2, 75.0, 2)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orders" LIMIT $1`)).
		WithArgs(10).
		WillReturnRows(rows)

	orders, total, err := repo.FindAll(1, 10, map[string]any{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if total != 2 {
		t.Fatalf("expected total 2, got %d", total)
	}

	if len(orders) != 2 {
		t.Fatalf("expected 2 orders, got %d", len(orders))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestFindOneByID(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewOrderRepository(db)

	rows := sqlmock.NewRows([]string{"id", "price", "user_id"}).
		AddRow(1, 120.00, 3)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orders" WHERE "orders"."id" = $1 ORDER BY "orders"."id" LIMIT $2`)).
		WithArgs(1, 1).
		WillReturnRows(rows)

	order, err := repo.FindOneByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if order.ID != 1 {
		t.Fatalf("expected order ID 1, got %d", order.ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
