package events

import (
	"github.com/rs/zerolog"
)

func (r APIGatewayV2HTTPRequest) MarshalZerologObject(e *zerolog.Event) {
	e.EmbedObject(r.RequestContext)
}

func (rc APIGatewayV2HTTPRequestContext) MarshalZerologObject(e *zerolog.Event) {
	e.Str("requestId", rc.RequestID).
		Str("accountId", rc.AccountID).
		Str("stage", rc.Stage).
		Str("routeKey", rc.RouteKey).
		Str("apiId", rc.APIID).
		Time("time", rc.TimeEpoch.Time).
		EmbedObject(rc.HTTP)
}

func (rch APIGatewayV2HTTPRequestContextHTTPDescription) MarshalZerologObject(e *zerolog.Event) {
	e.Str("method", rch.Method).
		Str("path", rch.Path).
		Str("protocol", rch.Protocol).
		Str("sourceIp", rch.SourceIP).
		Str("userAgent", rch.UserAgent)
}
