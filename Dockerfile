# This Dockerfile leverages multi-stage builds.
# See https://docs.docker.com/develop/develop-images/multistage-build/

# STAGE:
# Fetch summon

FROM ruby:3.0 as summon

RUN apt-get update && \
    apt-get install -y --no-install-recommends curl

# Install summon and summon-conjur
RUN curl -sSL https://raw.githubusercontent.com/cyberark/summon/master/install.sh \
      | env TMPDIR=$(mktemp -d) bash && \
    curl -sSL https://raw.githubusercontent.com/cyberark/summon-conjur/master/install.sh \
      | env TMPDIR=$(mktemp -d) bash

# STAGE:
# The 'maven' base is used to package the application
FROM maven:3.8.5-openjdk-11-slim as maven

WORKDIR /app

# To make best use of Docker caching this section only copies
# pom.xml and uses it to install all the dependencies required
# to package the application in the next section
COPY ./pom.xml ./pom.xml
RUN mvn verify clean --fail-never

# copy the source and package the application
COPY ./src ./src
RUN mvn package && cp target/petstore-*.jar app.jar

# STAGE:
# This base is used for the final image
# It extracts the packaged application from the previous stage
# and builds the final image
FROM openjdk:11-jdk-slim
LABEL org.opencontainers.image.authors="CyberArk"

COPY --from=summon /usr/local/lib/summon /usr/local/lib/summon
COPY --from=summon /usr/local/bin/summon /usr/local/bin/summon
COPY --from=maven /app/app.jar /app.jar

ENTRYPOINT [ "java", "-jar", "/app.jar"]
