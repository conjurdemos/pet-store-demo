# This Dockerfile leverages multi-stage builds.
# See https://docs.docker.com/develop/develop-images/multistage-build/

# STAGE:
# The 'maven' base is used to package the application
FROM maven:3.5.3-jdk-8-alpine as maven

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
FROM java:8-jre-alpine
MAINTAINER CyberArk

COPY --from=maven /app/app.jar /app.jar

ENTRYPOINT [ "java", "-jar", "/app.jar"]
