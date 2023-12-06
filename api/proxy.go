package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func ProxyHandler(targetHost string, beforeRequestFn func(req *http.Request),
	formatResponse func(resp *http.Response) error) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Director = func(request *http.Request) {
		targetQuery := url.RawQuery
		request.URL.Scheme = url.Scheme
		request.URL.Host = url.Host
		request.Host = url.Host
		request.URL.Path, request.URL.RawPath = joinURLPath(url, request.URL)

		if targetQuery == "" || request.URL.RawQuery == "" {
			request.URL.RawQuery = targetQuery + request.URL.RawQuery
		} else {
			request.URL.RawQuery = targetQuery + "&" + request.URL.RawQuery
		}
		if _, ok := request.Header["User-Agent"]; !ok {
			request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36")
			request.Header.Set("Content-Type", "application/json; charset=utf-8")
		}
		if beforeRequestFn != nil {
			beforeRequestFn(request)
		}
	}

	proxy.ModifyResponse = func(resp *http.Response) error {
		if formatResponse != nil {
			return formatResponse(resp)
		}
		return nil
	}

	return proxy, err
}

func JsonProxyRequestHandler(proxyFunc func(*gin.Context) (*httputil.ReverseProxy, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		proxy, err := proxyFunc(ctx)
		if err != nil {
			ctx.JSON(SERVER_ERROR.getHttpCode(), map[string]any{
				"code":    SERVER_ERROR,
				"message": err.Error(),
			})
			return
		}
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}