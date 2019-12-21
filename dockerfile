# Ubuntu LTS by default
FROM ubuntu:18.04
RUN apt-get update && apt-get install -y \
  golang-go \
  make
COPY . /joueur
RUN cd /joueur && make

# Your client image should now be ready to run on an arena.
# If you use require additional dependencies add them here as a RUN command.
# Otherwise, at runtime your clinet will NOT have internet access.
