# this is docker multistage build

FROM golang:1.14-alpine3.12 AS builder

LABEL stage="builder"

LABEL maintainer="narendran <narendev0610@gmail.com>"

ENV GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /productapi

# copying the all the codes
COPY . .

#  (optional) install any compiler-only dependencies
RUN apk add --no-cache gcc libc-dev  \
    &&  go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build . 


FROM alpine AS final

LABEL stage="final"

LABEL maintainer="narendran <narendev0610@gmail.com>"

WORKDIR /finalapp

# copying from builder the GO executable file
COPY --from=builder /productapi/go-micro-learning .

# app needs swagger.yaml to serve the docs properly on the /docs route
COPY --from=builder /productapi/swagger.yaml .

# app uses this port , so it is exposed.
EXPOSE 9090

#  executing the program upon start 
CMD [ "./go-micro-learning" ]