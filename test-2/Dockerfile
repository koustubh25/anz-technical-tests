############################
# STEP 1 build source code
############################
FROM golang:alpine as builder
LABEL authors="Koustubh"

ENV GO111MODULE=on

WORKDIR /app

# Create appuser
ENV USER=appuser
ENV UID=10001
 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

ADD . .   

RUN apk add jq make

RUN make all

# #################################
# # STEP 2 Copy executable and run 
# #################################
FROM scratch

COPY --from=builder /app/golang-version /.
COPY --from=builder /app/resource /resource
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Use an unprivileged user.
USER appuser:appuser

ENTRYPOINT ["/golang-version"]
EXPOSE 8000
