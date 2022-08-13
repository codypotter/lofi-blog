FROM golang:1.19-bullseye

# Set environment variable
ENV APP_NAME lofi-blog
ENV CMD_PATH main.go
 
# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
 
# Build application
RUN CGO_ENABLED=0  go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

EXPOSE 8080

CMD ./$APP_NAME