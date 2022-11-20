package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/exp/slog"
)

type Item struct {
	Name string
	Size int
}

func main() {
	start := time.Now()
	item := Item{
		Name: "item name",
		Size: 1111,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout))
	logger.Info("processed item", "name", item.Name, "item", item.Size, "duration", time.Since(start))
	logger.LogDepth(-1000, slog.InfoLevel, "processed item", "name", item.Name, "item", item.Size, "duration", time.Since(start))
	logger.Debug("debug!!!")

	// slog.LogAttrs(slog.InfoLevel, "processed item", slog.String("name", item.Name))

	slog.SetDefault(logger)
	log.Printf("processed %s of size %d in %s", item.Name, item.Size, time.Since(start))

	fmt.Println("------------------------------------")
	myLogger := slog.New(newMyHandler(os.Stdout))
	myLogger.Info("some INFO text")
	myLogger.Error("some Error text", errors.New("some error!"))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	logger := slog.Default().With("URL", r.URL)
	logger.Info("start")
	ctx := slog.NewContext(r.Context(), logger)
	r = r.WithContext(ctx)
	process(w, r)
}

func process(w http.ResponseWriter, r *http.Request) {
	logger := slog.FromContext(r.Context())
	logger.Debug("in process", "key", "some value")
	// ...
}

type MyLogLevel slog.Level

const MyLogLevelDetail MyLogLevel = 1

type MyHandler struct {
	writer io.Writer
}

func newMyHandler(writer io.Writer) *MyHandler {
	return &MyHandler{
		writer: writer,
	}
}

func (h *MyHandler) Enabled(level slog.Level) bool {
	return level >= slog.Level(MyLogLevelDetail)
}

func (h *MyHandler) Handle(r slog.Record) error {
	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("LEVEL=%s", r.Level.String()))
	buf.WriteByte(' ')
	buf.WriteString(fmt.Sprintf("msg: %s, err detail: %s\n", r.Message, "")) //エラーのattrsはどうやって出力するのが良いのか？
	_, err := h.writer.Write(buf.Bytes())
	return err
}

func (h *MyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return slog.Default().Handler()
}

func (h *MyHandler) WithGroup(name string) slog.Handler {
	return slog.Default().Handler()
}
