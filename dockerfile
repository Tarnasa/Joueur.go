FROM golang:1.15.5-buster
WORKDIR /client
COPY . .
RUN make
ENTRYPOINT ["joueur"]

# Your client image should now be ready to run on an arena.
# If you use require additional dependencies add them here as a RUN command.
# Otherwise, at runtime your clinet will NOT have internet access.
