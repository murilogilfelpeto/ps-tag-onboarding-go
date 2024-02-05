FROM golang:1.21-alpine

WORKDIR /app
COPY go.mod .
RUN go mod download

COPY . .

WORKDIR /app/cmd/ps-tag-onboarding
RUN go build -o ps-tag-onboarding

EXPOSE 8080
ENV ENVIRONMENT=docker-macos
CMD [ "./ps-tag-onboarding" ]