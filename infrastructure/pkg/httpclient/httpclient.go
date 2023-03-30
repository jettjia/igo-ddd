package httpclient

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"

	"github.com/jettjia/go-ddd-demo/global"
)

// HttpClient http请求
func HttpClient(ctx context.Context, apiUrl string, method string, reqParams interface{}, headers map[string]string, isDebug bool) (string, error) {
	var rest string

	gClient := g.Client()

	for k, v := range headers {
		gClient.SetHeader(k, v)
	}

	if isDebug {
		g.Log().Debug(ctx, "netClient apiUrl:", apiUrl)
		g.Log().Debug(ctx, "netClient method:", method)
		g.Log().Debug(ctx, "netClient params:", reqParams)
		g.Log().Debug(ctx, "netClient header:", headers)
		g.Log().Debug(ctx, "netClient ctx:", ctx)
	}

	if method == "GET" || method == "get" {
		r, err := gClient.Get(ctx, apiUrl)
		if err != nil {
			return rest, err
		}
		defer func(r *gclient.Response) {
			err := r.Close()
			if err != nil {
				global.GLog.Errorln("HttpClient:gclient:close err:", err.Error())
			}
		}(r)

		if r.StatusCode != 200 && r.StatusCode != 201 && r.StatusCode != 204 {
			return "", errors.New("HttpClient:get method error:" + gconv.String(r.StatusCode))
		}
		return r.ReadAllString(), nil
	}

	if method == "POST" || method == "post" {
		bytes, _ := json.Marshal(reqParams)
		r, err := gClient.Post(ctx, apiUrl, string(bytes))
		if err != nil {
			return rest, err
		}
		defer func(r *gclient.Response) {
			err := r.Close()
			if err != nil {
				global.GLog.Errorln("HttpClient:gclient:close err:", err.Error())
			}
		}(r)

		if r.StatusCode != 200 && r.StatusCode != 201 && r.StatusCode != 204 {
			return "", errors.New("HttpClient:post method error:" + gconv.String(r.StatusCode))
		}

		return r.ReadAllString(), nil
	}

	if method == "POSTFORM" || method == "postform" {
		r, err := gClient.SetContentType("application/x-www-form-urlencoded").
			Post(ctx, apiUrl, reqParams)

		if err != nil {
			return rest, err
		}
		defer func(r *gclient.Response) {
			err := r.Close()
			if err != nil {
				global.GLog.Errorln("HttpClient:gclient:close err:", err.Error())
			}
		}(r)

		if r.StatusCode != 200 && r.StatusCode != 201 && r.StatusCode != 204 {
			return "", errors.New("HttpClient:postform method error:" + gconv.String(r.StatusCode))
		}

		return r.ReadAllString(), nil
	}

	if method == "DELETE" || method == "delete" || method == "del" || method == "DEL" {
		r, err := gClient.Delete(ctx, apiUrl)
		if err != nil {
			return rest, err
		}
		defer func(r *gclient.Response) {
			err := r.Close()
			if err != nil {
				global.GLog.Errorln("HttpClient:gclient:close err:", err.Error())
			}
		}(r)

		if r.StatusCode != 200 && r.StatusCode != 201 && r.StatusCode != 204 {
			return "", errors.New("HttpClient:delete method error:" + gconv.String(r.StatusCode))
		}
		return r.ReadAllString(), nil
	}

	if method == "PUT" || method == "put" {
		bytes, _ := json.Marshal(reqParams)

		r, err := gClient.Put(ctx, apiUrl, bytes)
		if err != nil {
			return rest, err
		}
		defer func(r *gclient.Response) {
			err := r.Close()
			if err != nil {
				global.GLog.Errorln("HttpClient:gclient:close err:", err.Error())
			}
		}(r)

		if r.StatusCode != 200 && r.StatusCode != 201 && r.StatusCode != 204 {
			return "", errors.New("HttpClient:put method error:" + gconv.String(r.StatusCode))
		}

		return r.ReadAllString(), nil
	}

	return rest, nil
}
