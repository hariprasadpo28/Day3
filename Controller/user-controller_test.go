package Controller

//import "C"
import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	//Config.DB
	//router := gin.Default()
	//router.DELETE("/student-data/student/:id", DeleteUser)
	//w := httptest.NewRecorder()
	//req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/student-data/student/8", nil)
	//router.ServeHTTP(w, req)
	//DeleteUser("8")
	//

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/student-data/student/10", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch Request
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.Equal(t, "200 OK", resp.Status)
	respBody, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, `{"message":"Student deleted"}`, respBody)
}