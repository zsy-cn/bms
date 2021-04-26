package parser

import (
	"context"
	"encoding/json"
	"net/http"

	transport_http "github.com/go-kit/kit/transport/http"

	"github.com/zsy-cn/bms/protos"
)

func decodeHTTPRequest(request interface{}) func(context.Context, *http.Request) (interface{}, error) {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			// consul在请求健康检查接口时请求体为空, 会导致json的Decode()错误, 这里忽略这种错误
			if err.Error() != "EOF" {
				logger.Errorf("decode request data failed: %s", err.Error())
				return nil, err
			}
		}
		return request, nil
	}
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// StartHTTPTransport ...
// 由于上行数据解析服务只需要一个方法, 一个路由, 所以直接使用原生http库即可.
func StartHTTPTransport(srv Parser) {
	decodeHandler := transport_http.NewServer(
		MakeDecodeEndpoint(srv),
		decodeHTTPRequest(&protos.ParserHubUplinkMsg{}),
		encodeHTTPResponse,
	)
	healthCheckHandler := transport_http.NewServer(
		MakeHealthCheckEndpoint(srv),
		decodeHTTPRequest(&protos.Empty{}),
		encodeHTTPResponse,
	)
	http.Handle("/parse", decodeHandler)
	http.Handle("/health", healthCheckHandler)
}
