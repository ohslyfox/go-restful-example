package api

import (
	"io"
	"net/http"
	"testing"

	"github.com/ohslyfox/go-restful-example/app/db"

	"gorm.io/gorm"
)

var Database *gorm.DB

func TestApi(t *testing.T) {
	Setup()
	t.Run("is-db-empty", func(t *testing.T) {
		actual := IsGetResponseEmpty()
		if !actual {
			t.Errorf("Response is not empty")
		}
	})
}

func Setup() {
	Database = db.GetDatabaseInstance()
	db.DeleteAllData(Database)
}

func IsGetResponseEmpty() bool {
	resp, err := http.Get("127.0.0.1:3000/books")
	if err != nil {
		return false;
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false;
	}

	if len(body) > 0 {
		return false;
	}
	return true;
}
