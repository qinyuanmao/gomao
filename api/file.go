package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qinyuanmao/gomao/logger"
)

type FileHandler func(ctx *Context) (resultCode ResultCode, message string, filePath, fileName, contentType string)

func FileApi(hander FileHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code, message, filePath, fileName, contentType := hander(&Context{ctx})
		if code != SUCCESS {
			sendDingTalk(ctx.Request.URL.String(), message, code.getHttpCode())
			ctx.String(code.getHttpCode(), message)
			return
		}
		file, err := os.Open(filePath)
		if err != nil {
			sendDingTalk(ctx.Request.URL.String(), fmt.Sprintf("文件打开异常：%s", err.Error()), http.StatusInternalServerError)
			ctx.String(http.StatusInternalServerError, message)
			return
		}
		sendFile(ctx.Writer, ctx.Request, file, contentType, fileName)
	}
}

func sendFile(writer http.ResponseWriter, request *http.Request, f *os.File, contentType, fileName string) {
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		http.NotFound(writer, request)
		return
	}
	writer.Header().Add("Accept-Ranges", "bytes")
	writer.Header().Add("Content-Type", contentType)
	writer.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))
	writer.Header().Set("Accept-Length", fmt.Sprintf("%d", info.Size()))
	writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	var start, end int64
	if r := request.Header.Get("Range"); r != "" {
		if strings.Contains(r, "bytes=") && strings.Contains(r, "-") {
			fmt.Sscanf(r, "bytes=%d-%d", &start, &end)
			if end == 0 {
				end = info.Size() - 1
			}
			if start > end || start < 0 || end < 0 || end >= info.Size() {
				writer.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
				return
			}
			writer.Header().Add("Content-Length", strconv.FormatInt(end-start+1, 10))
			writer.Header().Add("Content-Range", fmt.Sprintf("bytes %v-%v/%v", start, end, info.Size()))
			writer.WriteHeader(http.StatusPartialContent)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		writer.Header().Add("Content-Length", strconv.FormatInt(info.Size(), 10))
		start = 0
		end = info.Size() - 1
	}
	_, err = f.Seek(start, 0)
	if err != nil {
		sendDingTalk(request.URL.String(), fmt.Sprintf("文件读取异常：%s", err.Error()), http.StatusInternalServerError)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	n := 512
	buf := make([]byte, n)
	for {
		if end-start+1 < int64(n) {
			n = int(end - start + 1)
		}
		_, err := f.Read(buf[:n])
		if err != nil {
			if err != io.EOF {
				logger.Errorf("read file error: %v", err)
			}
			return
		}
		err = nil
		_, err = writer.Write(buf[:n])
		if err != nil {
			return
		}
		start += int64(n)
		if start >= end+1 {
			return
		}
	}
}
