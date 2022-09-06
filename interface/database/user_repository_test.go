package database_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"clean-architecture/domain"
	"clean-architecture/infra"
	"clean-architecture/interface/database"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dummyGormHandler(conn *gorm.DB) database.GormHandler {
	gormHandler := new(infra.GormHandler)
	gormHandler.DB = conn
	return gormHandler
}

func newDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	// sqlmockを初期化
	mockDB, mock, err := sqlmock.New()

	if err != nil {
		return nil, mock, err
	}

	// DBのコネクションを初期化
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      mockDB,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{},
	)

	return db, mock, err
}

func TestFindAll(t *testing.T) {
	mockDB, mock, err := newDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mockDB: %v", err)
	}
	//期待されるカラム
	expectedColumns := []string{
		"article_id",
		"user_id",
		"title",
		"content",
		"created_at",
		"updated_at",
	}
	// 期待される結果
	expect := []domain.User{{
		UserId:       "userId1",
		Name:         "name1",
		FavoriteFood: "favoriteFood1",
		Birthday:     time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}}

	// 期待されるSQL
	query := "SELECT * FROM `users`"

	//DBの振る舞いを定義する
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(
			sqlmock.NewRows(expectedColumns).AddRow(expect[0].UserId, expect[0].Name, expect[0].FavoriteFood, expect[0].Birthday, expect[0].CreatedAt, expect[0].UpdatedAt))
	// repository 初期化
	repo := &database.UserRepository{GormHandler: dummyGormHandler(mockDB)}
	articles, _ := repo.FindAll()
	fmt.Println(articles)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test FindALL Article: %v", err)
	}
}
