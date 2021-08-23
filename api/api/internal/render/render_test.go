package render

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSuccess(t *testing.T) {
	resp := Success(nil)

	if resp.Status != "success" {
		t.Fatalf("Msg must be 'success'. Get=%s", resp.Status)
	}
	if resp.ErrorCode != 0 {
		t.Fatalf("SubCode should be 0. Get=%d", resp.ErrorCode)
	}
	if resp.ErrorMsg != "" {
		t.Fatalf("SubMsg must be empty. Get=%s", resp.ErrorMsg)
	}
	if _, ok := resp.Data.(gin.H); !ok {
		t.Fatalf("Data must be empty map. Get=%T", resp.Data)
	}
}
