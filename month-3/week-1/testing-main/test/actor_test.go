package order_service

import (
	"crud/models"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

var k int64

func TestActor(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			createActor(t)
			// deleteActor(t, id)
		}()
		k++
	}

	wg.Wait()

	fmt.Println("s: ", k)
}

func createActor(t *testing.T) string {
	response := &models.Actor{}

	request := &models.CreateActor{
		First_name: faker.FirstName(),
		Last_name:  faker.LastName(),
	}

	resp, err := PerformRequest(http.MethodPost, "/actor", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.Id
}

func updateActor(t *testing.T, id string) string {
	response := &models.Actor{}
	request := &models.UpdateActor{
		First_name: faker.FirstName(),
		Last_name:  faker.LastName(),
	}

	resp, err := PerformRequest(http.MethodPut, "/actor/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.Id
}

func deleteActor(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/actor/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
