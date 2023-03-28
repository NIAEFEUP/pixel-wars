FROM alpine:3.17.2
WORKDIR /pixel-wars
RUN apk update
RUN apk add nodejs npm go
ADD . /pixel-wars
WORKDIR /pixel-wars/frontend
RUN npm i
RUN npx vite build
WORKDIR /pixel-wars/backend
RUN go build

EXPOSE 80

ENTRYPOINT  ./backend-nixel-wars --prod
