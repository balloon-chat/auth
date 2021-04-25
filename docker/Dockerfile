FROM golang:latest
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
ENV GOOGLE_APPLICATION_CREDENTIALS=/go/src/app/secret/balloon-6bad2-firebase-adminsdk-q8zwv-fb04c9c9e8.json
CMD ["go", "run", "cmd/main.go"]
