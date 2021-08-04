package LinearSort

import (
	"testing"
	"unsafe"
)

func TestCountingSort(t *testing.T) {
	arr := []int{5, 4}
	CountingSort(arr, len(arr))
	t.Log(arr)

	arr = []int{5, 4, 3, 2, 1}
	CountingSort(arr, len(arr))
	t.Log(arr)
	s := &DealPo{}
	t.Logf("s=%+v", s)
	t1 := unsafe.Pointer(s)
	t.Logf("t1=%+v", t1)
}

type DealPo struct {
	M_strDealId     string `json:"DealId"`
	M_lDealId64     int64  `json:"DealId64"`
	M_lRecvFeeId    int64  `json:"RecvFeeId"`
	M_lBuyerId      int64  `json:"BuyerId"`
	M_strBuyerNick  string `json:"BuyerNick"`
	M_lSellerId     int64  `json:"SellerId"`
	M_lSellerCorpId int64  `json:"SellerCorpId"`
}
