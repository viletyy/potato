/*
 * @Date: 2021-06-13 23:25:52
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-14 20:42:40
 * @FilePath: /potato/initialize/tracer.go
 */
package initialize

import (
	"github.com/opentracing/opentracing-go"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/pkg/tracer"
)

func Tracer() opentracing.Tracer {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		global.GO_CONFIG.App.Name,
		global.GO_CONFIG.Server.TracerHostPort,
	)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("tracer.NewJaegerTracer err: %v", err)
	}

	return jaegerTracer
}
