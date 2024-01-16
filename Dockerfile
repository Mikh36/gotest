FROM golang:1.21 as build
LABEL stage=intermediate
WORKDIR /app
COPY . .
RUN make build
RUN ls -lah /app/bin/

FROM scratch as scratch
COPY --from=build /app/bin/testapp /bin/testapp
EXPOSE 8080
CMD ["testapp"]