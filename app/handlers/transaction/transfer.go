package transaction

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"mypayment/business/core/transaction"
	"mypayment/foundation/web"
)

func (h Handler) Transfer(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var rt transferReq
	if err := web.Decode(r, &rt); err != nil {
		return fmt.Errorf("decoding: %w: %v", errBadRequest, err)
	}

	result, err := h.s.InsertTransfer(ctx, toTransferModel(rt))
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, present(result), http.StatusCreated)
}

func (h Handler) UpdateStatus(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var req struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}

	if err := web.Decode(r, &req); err != nil {
		return fmt.Errorf("decoding: %w: %v", errBadRequest, err)
	}

	result, err := h.s.UpdateStatus(ctx, req.ID, req.Status)
	if err != nil {
		return err
	}
	return web.Respond(ctx, w, updateResp(result), http.StatusOK)
}

func toTransferModel(rt transferReq) transaction.RawTransfer {
	return transaction.RawTransfer{
		SenderName:         rt.SenderName,
		SenderBankNumber:   rt.SenderBankNumber,
		SenderBank:         rt.SenderBank,
		ReceiverName:       rt.ReceiverName,
		ReceiverBankNumber: rt.ReceiverBankNumber,
		ReceiverBank:       rt.ReceiverBank,
		Amount:             rt.Amount,
		Currency:           rt.Currency,
		Description:        rt.Description,
	}
}

func present(val transaction.Transfer) transferResponse {
	return transferResponse{
		ID:                 val.ID,
		SenderName:         val.SenderName,
		SenderBankNumber:   val.SenderBankNumber,
		SenderBank:         val.SenderBank,
		ReceiverName:       val.ReceiverName,
		ReceiverBankNumber: val.ReceiverBankNumber,
		ReceiverBank:       val.ReceiverBank,
		Amount:             val.Amount,
		Currency:           val.Currency,
		Status:             val.Status,
		Description:        val.Description,
		Created:            val.Created,
		Updated:            val.Updated,
	}
}

type updateResponse struct {
	ID      string    `json:"id"`
	Status  string    `json:"status"`
	Updated time.Time `json:"updated"`
}

func updateResp(val transaction.StatusUpdate) updateResponse {
	return updateResponse{
		ID:      val.ID,
		Status:  val.Status,
		Updated: val.Updated,
	}
}
