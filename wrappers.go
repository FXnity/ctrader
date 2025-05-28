package ctrader

import (
	"context"
	"github.com/fxnity/ctrader/openapi"
)

func (c *Client) AccountList(accessToken string) ([]*openapi.ProtoOACtidTraderAccount, error) {
	accountListReq := &openapi.ProtoOAGetAccountListByAccessTokenReq{
		AccessToken: &accessToken,
	}
	al, err := Command[*openapi.ProtoOAGetAccountListByAccessTokenReq, *openapi.ProtoOAGetAccountListByAccessTokenRes](context.Background(), c, accountListReq)
	if err != nil {
		return nil, err
	}
	return al.CtidTraderAccount, nil
}

func (c *Client) AccountAuthorize(accessToken string, CtidTraderAccountId int64) error {
	req := &openapi.ProtoOAAccountAuthReq{
		AccessToken:         &accessToken,
		CtidTraderAccountId: &CtidTraderAccountId,
	}
	_, err := Command[*openapi.ProtoOAAccountAuthReq, *openapi.ProtoOAAccountAuthRes](context.Background(), c, req)
	return err
}

func (c *Client) SymbolsList(CtidTraderAccountId int64) ([]*openapi.ProtoOALightSymbol, error) {
	reqSymbolList := &openapi.ProtoOASymbolsListReq{
		CtidTraderAccountId: &CtidTraderAccountId,
	}
	respSymbolList, err := Command[*openapi.ProtoOASymbolsListReq, *openapi.ProtoOASymbolsListRes](
		context.Background(), c, reqSymbolList,
	)
	if err != nil {
		return nil, err
	}
	return respSymbolList.GetSymbol(), nil
}

func (c *Client) SubscribeSpots(CtidTraderAccountId int64, symbols []int64, withTimeStamps bool) error {

	subscribeRequest := &openapi.ProtoOASubscribeSpotsReq{
		CtidTraderAccountId:      &CtidTraderAccountId,
		SymbolId:                 symbols,
		SubscribeToSpotTimestamp: &withTimeStamps,
	}
	_, err := Command[*openapi.ProtoOASubscribeSpotsReq, *openapi.ProtoOASubscribeSpotsRes](context.Background(), c, subscribeRequest)
	return err
}
