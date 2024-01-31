FROM golang:1.21-alpine

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o ps-tag-onboarding-go

EXPOSE 8080
ENV ENVIRONMENT=docker-macos
CMD [ "./ps-tag-onboarding-go" ]