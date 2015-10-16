package log

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const fatalEnvKey = "TEST_FATAL"

var dummy = &bytes.Buffer{}

func TestMain(m *testing.M) {

	logger.Out = dummy
	logger.Level = logrus.DebugLevel
	logger.Formatter = &logrus.JSONFormatter{}

	if _, exist := os.LookupEnv(fatalEnvKey); exist {
		Fatal("fatal test message", errors.New("dummy"))
	}

	os.Exit(m.Run())
}

func TestAddMetadata(t *testing.T) {

	dummy.Reset()

	k, v := "hoge", "fuga"
	AddMetadata(logrus.Fields{k: v})

	Info("test")
	time.Sleep(time.Millisecond)

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(dummy).Decode(&m))
	assert.Equal(t, v, m[k])
}

func TestSetLogger(t *testing.T) {

	tmp := logger
	assert.Equal(t, logrus.DebugLevel, logger.Level)

	SetLogger(&logrus.Logger{Level: logrus.InfoLevel})
	assert.Equal(t, logrus.InfoLevel, logger.Level)

	SetLogger(tmp)
	assert.Equal(t, logrus.DebugLevel, logger.Level)
}

func TestAddHooks(t *testing.T) {

	AddHooks()
	assert.Empty(t, logger.Hooks)

	h := &TestHook{}
	AddHooks(h)
	assert.Equal(t, len(h.Levels()), len(logger.Hooks))
}

func TestDebug(t *testing.T) {

	// only message

	dummy.Reset()

	msg := "debug test message"
	Debug(msg)
	time.Sleep(time.Millisecond)

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(dummy).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.DebugLevel.String(), m["level"])
	assert.Nil(t, m["id"])

	// with fields

	dummy.Reset()

	id := "vytxeTZskVKR7C7WgdSP3d"
	Debug(msg, logrus.Fields{"id": id})
	time.Sleep(time.Millisecond)

	m = map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(dummy).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.DebugLevel.String(), m["level"])
	assert.Equal(t, id, m["id"])
}

func TestInfo(t *testing.T) {

	// only message

	dummy.Reset()

	msg := "info test message"
	Info(msg)
	time.Sleep(time.Millisecond)

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(dummy).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.InfoLevel.String(), m["level"])
	assert.Nil(t, m["id"])

	// with fields

	dummy.Reset()

	id := "vytxeTZskVKR7C7WgdSP3d"
	Info(msg, logrus.Fields{"id": id})
	time.Sleep(time.Millisecond)

	m = map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(dummy).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.InfoLevel.String(), m["level"])
	assert.Equal(t, id, m["id"])
}

func TestWarn(t *testing.T) {

	// only message

	dummy.Reset()

	msg := "warn test message"
	Warn(msg)
	time.Sleep(time.Millisecond)

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(dummy).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.WarnLevel.String(), m["level"])
	assert.Nil(t, m["id"])

	// with fields

	dummy.Reset()

	id := "vytxeTZskVKR7C7WgdSP3d"
	Warn(msg, logrus.Fields{"id": id})
	time.Sleep(time.Millisecond)

	m = map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(dummy).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.WarnLevel.String(), m["level"])
	assert.Equal(t, id, m["id"])
}

func TestError(t *testing.T) {

	// only message

	dummy.Reset()

	msg := "error test message"
	err := errors.New("dummy")
	Error(msg, err)
	time.Sleep(time.Millisecond)

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(dummy).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.ErrorLevel.String(), m["level"])
	assert.Equal(t, err.Error(), m[logrus.ErrorKey])
	assert.Nil(t, m["id"])

	// with fields

	dummy.Reset()

	id := "vytxeTZskVKR7C7WgdSP3d"
	Error(msg, err, logrus.Fields{"id": id})
	time.Sleep(time.Millisecond)

	m = map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(dummy).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.ErrorLevel.String(), m["level"])
	assert.Equal(t, err.Error(), m[logrus.ErrorKey])
	assert.Equal(t, id, m["id"])
}

func TestFatal(t *testing.T) {

	cmd := exec.Command(os.Args[0])
	cmd.Env = append(os.Environ(), fatalEnvKey+"=1")
	assert.Error(t, cmd.Run())
}

func TestPanic(t *testing.T) {

	// only message

	dummy.Reset()

	msg := "panic test message"
	err := errors.New("dummy")
	assert.Panics(t, func() { Panic(msg, err) })

	m := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(dummy).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.PanicLevel.String(), m["level"])
	assert.Equal(t, err.Error(), m[logrus.ErrorKey])
	assert.Nil(t, m["id"])

	// with fields

	dummy.Reset()

	id := "vytxeTZskVKR7C7WgdSP3d"
	assert.Panics(t, func() { Panic(msg, err, logrus.Fields{"id": id}) })

	m = map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(dummy).Decode(&m))

	assert.Equal(t, msg, m["msg"])
	assert.Equal(t, logrus.PanicLevel.String(), m["level"])
	assert.Equal(t, err.Error(), m[logrus.ErrorKey])
	assert.Equal(t, id, m["id"])
}

type TestHook struct {
	Fired bool
}

func (hook *TestHook) Fire(entry *logrus.Entry) error {
	hook.Fired = true
	return nil
}

func (hook *TestHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
