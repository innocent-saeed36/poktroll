package uzap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pokt-network/poktroll/pkg/ulogger"
)

var _ ulogger.UniversalLogger = (*zapULogger)(nil)

type zapULogger struct {
	// NB: Default (0) is Info.
	level         zapcore.Level
	writeSyncer   zapcore.WriteSyncer
	encoderConfig zapcore.EncoderConfig
	encoder       zapcore.Encoder
	core          zapcore.Core
	logger        *zap.Logger
}

func NewUniversalLogger(
	opts ...ulogger.LoggerOption,
) ulogger.UniversalLogger {
	ze := &zapULogger{}

	for _, opt := range opts {
		opt(ze)
	}

	ze.buildLoggerAndSetDefaults()

	return ze
}

func (za *zapULogger) Debug() ulogger.Event {
	return newEvent(za.logger, zapcore.DebugLevel)
}

func (za *zapULogger) Info() ulogger.Event {
	return newEvent(za.logger, zapcore.InfoLevel)
}

func (za *zapULogger) Warn() ulogger.Event {
	return newEvent(za.logger, zapcore.WarnLevel)
}

func (za *zapULogger) Error() ulogger.Event {
	return newEvent(za.logger, zapcore.ErrorLevel)
}

// TODO_IN_THIS_COMMIT: Implement Fatal & Panic ?

func (za *zapULogger) buildLoggerAndSetDefaults() {
	if za.writeSyncer == nil {
		za.writeSyncer = zapcore.AddSync(os.Stderr)
	}

	if za.logger == nil {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoder := zapcore.NewJSONEncoder(encoderConfig)
		core := zapcore.NewCore(encoder, za.writeSyncer, za.level)

		za.logger = zap.New(core)
	}

}