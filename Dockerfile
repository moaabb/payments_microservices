FROM gcr.io/distroless/static-debian11
WORKDIR /app
ADD output .
EXPOSE 8080
ENTRYPOINT [ "/app/customer_svc" ]
