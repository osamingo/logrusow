package log

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	al = &bytes.Buffer{}
	el = &bytes.Buffer{}
)

func reset() {
	al.Reset()
	el.Reset()
}

func TestMain(m *testing.M) {

	// set mock writers
	appLogger.Out = al
	errLogger.Out = el

	os.Exit(m.Run())
}

func TestAddMetaInfo(t *testing.T) {

	reset()

	k, v := "hoge", "fuga"
	AddMetaInfo(logrus.Fields{k: v})

	Info("test")
	time.Sleep(time.Millisecond)

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(al).Decode(&m))
	assert.Equal(t, v, m[k])
}

func TestSetAppLogger(t *testing.T) {

	tmp := appLogger
	assert.Equal(t, logrus.DebugLevel, appLogger.Level)

	SetAppLogger(&logrus.Logger{Level: logrus.InfoLevel})
	assert.Equal(t, logrus.InfoLevel, appLogger.Level)

	SetAppLogger(tmp)
	assert.Equal(t, logrus.DebugLevel, appLogger.Level)
}

func TestSetErrLogger(t *testing.T) {

	tmp := errLogger
	assert.Equal(t, logrus.ErrorLevel, errLogger.Level)

	SetErrLogger(&logrus.Logger{Level: logrus.InfoLevel})
	assert.Equal(t, logrus.InfoLevel, errLogger.Level)

	SetErrLogger(tmp)
	assert.Equal(t, logrus.ErrorLevel, errLogger.Level)
}

func TestDebug(t *testing.T) {

	// only message

	reset()

	msg := "debug test message"
	Debug(msg)
	time.Sleep(time.Millisecond)

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(al).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.DebugLevel.String(), m["level"])
	assert.Nil(t, m["id"])

	// with fields

	reset()

	id := "vytxeTZskVKR7C7WgdSP3d"
	Debug(msg, logrus.Fields{"id": id})
	time.Sleep(time.Millisecond)

	m = map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(al).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.DebugLevel.String(), m["level"])
	assert.Equal(t, id, m["id"])
}

func TestInfo(t *testing.T) {

	// only message

	reset()

	msg := "info test message"
	Info(msg)
	time.Sleep(time.Millisecond)

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(al).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.InfoLevel.String(), m["level"])
	assert.Nil(t, m["id"])

	// with fields

	reset()

	id := "vytxeTZskVKR7C7WgdSP3d"
	Info(msg, logrus.Fields{"id": id})
	time.Sleep(time.Millisecond)

	m = map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(al).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.InfoLevel.String(), m["level"])
	assert.Equal(t, id, m["id"])
}

func TestWarn(t *testing.T) {

	// only message

	reset()

	msg := "warn test message"
	Warn(msg)
	time.Sleep(time.Millisecond)

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(al).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.WarnLevel.String(), m["level"])
	assert.Nil(t, m["id"])

	// with fields

	reset()

	id := "vytxeTZskVKR7C7WgdSP3d"
	Warn(msg, logrus.Fields{"id": id})
	time.Sleep(time.Millisecond)

	m = map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(al).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.WarnLevel.String(), m["level"])
	assert.Equal(t, id, m["id"])
}

func TestError(t *testing.T) {

	// only message

	reset()

	msg := "error test message"
	err := errors.New("dummy")
	Error(msg, err)
	time.Sleep(time.Millisecond)

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(el).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.ErrorLevel.String(), m["level"])
	assert.Equal(t, err.Error(), m[logrus.ErrorKey])
	assert.Nil(t, m["id"])

	// with fields

	reset()

	id := "vytxeTZskVKR7C7WgdSP3d"
	Error(msg, err, logrus.Fields{"id": id})
	time.Sleep(time.Millisecond)

	m = map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(el).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.ErrorLevel.String(), m["level"])
	assert.Equal(t, err.Error(), m[logrus.ErrorKey])
	assert.Equal(t, id, m["id"])
}

func TestPanic(t *testing.T) {

	// only message

	reset()

	msg := "panic test message"
	err := errors.New("dummy")
	assert.Panics(t, func() { Panic(msg, err) })

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(el).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.PanicLevel.String(), m["level"])
	assert.Equal(t, err.Error(), m[logrus.ErrorKey])
	assert.Nil(t, m["id"])

	// with fields

	reset()

	id := "vytxeTZskVKR7C7WgdSP3d"
	assert.Panics(t, func() { Panic(msg, err, logrus.Fields{"id": id}) })

	m = map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(el).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.PanicLevel.String(), m["level"])
	assert.Equal(t, err.Error(), m[logrus.ErrorKey])
	assert.Equal(t, id, m["id"])
}
