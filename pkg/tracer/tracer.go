/*
 * @Date: 2021-06-13 23:16:55
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-13 23:24:34
 * @FilePath: /potato/pkg/tracer/tracer.go
 */
package tracer

import (
	"io"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	// 设置应用的基本信息
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		}, // 固定采样、对所有数据都进行采样
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		}, // 是否启用LoggingReporter、刷新缓冲区都频率、上报的Agent地址
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
