import { createLogger, transports, format } from "winston";

const logFormat = format.printf(({ level, message, timestamp }) => {
  return `${timestamp} [${level}] ${message}`;
});

const winstonFormat = format.combine(format.timestamp(), logFormat);

export const logger = createLogger({
  transports: [
    new transports.Console({
      format: winstonFormat,
    }),
    new transports.File({
      filename: "error.log",
      level: "error",
      format: winstonFormat,
    }),
    new transports.File({
      filename: "combined.log",
      format: winstonFormat,
    }),
  ],
});
