FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["sh", "-c", "echo DATABASE_URL=$DATABASE_URL && ./app"]

