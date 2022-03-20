import winston from "winston";

export const logger = winston.createLogger({
  transports: [
    new winston.transports.Console({ format: winston.format.simple() }),
    new winston.transports.File({
      filename: "error.log",
      level: "error",
      format: winston.format.simple(),
    }),
    new winston.transports.File({
      filename: "combined.log",
      format: winston.format.simple(),
    }),
  ],
});
