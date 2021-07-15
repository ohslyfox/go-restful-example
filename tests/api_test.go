package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/ohslyfox/go-restful-example/app/db"
)

type TestData struct {
	Data []db.Book `json:"data"`
}

var Data *TestData

func TestApi(t *testing.T) {
	Setup()
	t.Run("delete-all", func(t *testing.T) {
		ErrorNotEqual(t, DeleteAll(), 204)
	})
	t.Run("is-response-empty", func(t *testing.T) {
		actual := IsGetResponseEmpty()
		if !actual {
			t.Errorf("Response is not empty")
		}
	})
	t.Run("list-all", func(t *testing.T) {
		ErrorNotEqual(t, ListAll(), 200)
	})
	t.Run("list-one-non-existing", func(t *testing.T) {
		ErrorNotEqual(t, ListOne(1), 404)
	})
	t.Run("insert-zeroth-record", func(t *testing.T) {
		ErrorNotEqual(t, InsertDataItem(0), 201)
	})
	t.Run("list-one", func(t *testing.T) {
		ErrorNotEqual(t, ListOne(1), 200)
	})
	t.Run("insert-duplicate-key-record", func(t *testing.T) {
		ErrorNotEqual(t, InsertDataItem(1), 409)
	})
	t.Run("insert-non-existing-record", func(t *testing.T) {
		ErrorNotEqual(t, InsertDataItem(-1), 400)
	})
	t.Run("insert-bad-record", func(t *testing.T) {
		ErrorNotEqual(t, InsertDataItem(2), 400)
	})
	t.Run("update-record", func(t *testing.T) {
		ErrorNotEqual(t, UpdateItem(1, 1), 200)
	})
	t.Run("update-non-existing-record", func(t *testing.T) {
		ErrorNotEqual(t, UpdateItem(2, 1), 404)
	})
	t.Run("update-with-bad-record", func(t *testing.T) {
		ErrorNotEqual(t, UpdateItem(1, 2), 400)
	})
	t.Run("checkout", func(t *testing.T) {
		ErrorNotEqual(t, CheckInOut(1, "checkout"), 200)
	})
	t.Run("checkout-non-existing", func(t *testing.T) {
		ErrorNotEqual(t, CheckInOut(2, "checkout"), 404)
	})
	t.Run("checkin", func(t *testing.T) {
		ErrorNotEqual(t, CheckInOut(1, "checkin"), 200)
	})
	t.Run("checkin-non-existing", func(t *testing.T) {
		ErrorNotEqual(t, CheckInOut(2, "checkin"), 404)
	})
	t.Run("delete-one", func(t *testing.T) {
		ErrorNotEqual(t, DeleteOne(1), 204)
	})
	t.Run("delete-one-non-existing", func(t *testing.T) {
		ErrorNotEqual(t, DeleteOne(1), 404)
	})
	t.Run("delete-all", func(t *testing.T) {
		ErrorNotEqual(t, DeleteAll(), 204)
	})
}

func ErrorNotEqual(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf(fmt.Sprintf("Received response code %d but expected %d", actual, expected))
	}
}

func Setup() {
	jsonFile, err := os.Open("test-data.json")
	if err != nil {
		panic("Missing test data")
	}
	bytes, _ := io.ReadAll(jsonFile)
	json.Unmarshal(bytes, &Data)
}

func IsGetResponseEmpty() bool {
	resp, err := http.Get("http://localhost:3000/books")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	if len(body) > 2 {
		return false
	}
	return true
}

func ListAll() int {
	res:= -1
	resp, _ := http.Get("http://localhost:3000/books")
	if resp != nil {
		res = resp.StatusCode
	}
	defer resp.Body.Close()
	return res
}

func ListOne(id int) int {
	res:= -1
	resp, _ := http.Get(fmt.Sprintf("http://localhost:3000/books/%d", id))
	if resp != nil {
		res = resp.StatusCode
	}
	defer resp.Body.Close()
	return res
}

func InsertDataItem(idx int) int {
	res := -1
	var resp *http.Response
	if idx < 0 {
		resp, _ = http.Post("http://localhost:3000/books", "application/json", nil)
	} else {
		testData, _ := json.Marshal(Data.Data[idx])
		resp, _ = http.Post("http://localhost:3000/books", "application/json", bytes.NewBuffer(testData))
	}
	if resp != nil {
		res = resp.StatusCode
	}
	defer resp.Body.Close()
	return res
}

func UpdateItem(findId int, dataIdx int) int {
	res := -1
	client := &http.Client{}
	var req *http.Request
	if findId < 0 {
		req, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:3000/books/%d", findId), nil)
	} else {
		testData, _ := json.Marshal(Data.Data[dataIdx])
		req, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:3000/books/%d", findId), bytes.NewBuffer(testData))
	}
	resp, _ := client.Do(req)
	if req != nil {
		res = resp.StatusCode
	}
	defer resp.Body.Close()
	return res
}

func DeleteAll() int {
	res := -1
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:3000/books", nil)
	resp, _ := client.Do(req)
	if resp != nil {
		res = resp.StatusCode
	}
	defer resp.Body.Close()
	return res
}

func DeleteOne(idx int) int {
	res := -1
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:3000/books/%d", idx), nil)
	resp, _ := client.Do(req)
	if resp != nil {
		res = resp.StatusCode
	}
	defer resp.Body.Close()
	return res
}

func CheckInOut(idx int, checkInOrOut string) int {
	res := -1
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:3000/books/%d/%s", idx, checkInOrOut), nil)
	resp, _ := client.Do(req)
	if resp != nil {
		res = resp.StatusCode
	}
	defer resp.Body.Close()
	return res
}
