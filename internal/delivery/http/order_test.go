package http

import (
	"encoding/json"
	"fmt"
	"github.com/Owner-maker/nats-learning/internal/models"
	mockService "github.com/Owner-maker/nats-learning/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

const orderJsonOutput = `{"order_uid":"b563feb7b2b84b6test","track_number":"WBILMTESTTRACK","entry":"WBIL","locale":"en","internal_signature":"","customer_id":"test","delivery_service":"meest","shardkey":"9","sm_id":99,"date_created":"2021-11-26T06:22:19Z","oof_shard":"1","delivery":{"name":"TestTestov","phone":"+9720000000","zip":"2639809","city":"KiryatMozkin","address":"PloshadMira15","region":"Kraiot","email":"test@gmail.com"},"payment":{"transaction":"b563feb7b2b84b6test","request_id":"","currency":"USD","provider":"wbpay","amount":1817,"payment_dt":1637907727,"bank":"alpha","delivery_cost":1500,"goods_total":317,"custom_fee":0},"items":[{"chrt_id":9934930,"track_number":"WBILMTESTTRACK","price":453,"rid":"ab4219087a764ae0btest","name":"Mascaras","sale":30,"size":"0","total_price":317,"nm_id":2389212,"brand":"VivienneSabo","status":202}]}`

func TestHandler_getOrderById(t *testing.T) {
	type mockBehavior func(s *mockService.MockOrder, uid string)
	var orderOutput models.Order

	err := json.Unmarshal([]byte(orderJsonOutput), &orderOutput)
	if err != nil {
		return
	}

	tests := []struct {
		name                 string
		inputUid             string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "Ok",
			inputUid: "b563feb7b2b84b6test",
			mockBehavior: func(s *mockService.MockOrder, uid string) {
				s.EXPECT().GetCachedOrder(uid).Return(orderOutput, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: orderJsonOutput,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			service := mockService.NewMockOrder(c)
			testCase.mockBehavior(service, testCase.inputUid)

			handler := NewHandler(service)

			// test server
			r := gin.New()
			r.GET("/order/:uid", handler.GetOrderById)

			//test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/order/%s", testCase.inputUid), nil)

			//perform request
			r.ServeHTTP(w, req)

			//assert result
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
