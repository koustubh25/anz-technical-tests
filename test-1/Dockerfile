############################
# STEP 1 build source code
############################
FROM golang:alpine as builder
LABEL authors="Koustubh"

ENV GO111MODULE=on

WORKDIR /app

# Ca-certificates is required to call HTTPS endpoints.
RUN apk update 
RUN apk add --no-cache ca-certificates && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

ADD . .    

RUN CGO_ENABLED=0  GOOS=linux go build -o golang-test .

# #################################
# # STEP 2 Copy executable and run 
# #################################
FROM scratch

COPY --from=builder /app/golang-test /.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Use an unprivileged user.
USER appuser:appuser

ENTRYPOINT ["/golang-test"]
EXPOSE 8000
