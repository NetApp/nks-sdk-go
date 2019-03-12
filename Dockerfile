FROM scratch
EXPOSE 8080
ENTRYPOINT ["/nks-sdk-go"]
COPY ./bin/ /