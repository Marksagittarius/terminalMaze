package logger

type LoggerPrinterMiddleware func(*Logger, string) string