package anticaptcha

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestClient_GetBalance(t *testing.T) {
	c := Client{
		ClientKey: os.Getenv("API_KEY"),
	}
	res, err := c.GetBalance()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.Balance)
}

func TestClient_GetQueueStats(t *testing.T) {
	queues := []int{1, 2, 5, 6, 7, 10}
	c := Client{}
	for _, queueId := range queues {
		res, err := c.GetQueueStats(queueId)
		if err != nil {
			t.Fatalf("queueId: %d %s", queueId, err)
		}
		resJson, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("queueId: %d error marshal %s", queueId, err)
		}
		t.Logf("%d - %s", queueId, string(resJson))
	}
}
