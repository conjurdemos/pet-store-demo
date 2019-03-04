FROM maven:3.5.3-jdk-8-alpine as maven

WORKDIR /app
COPY ./pom.xml ./pom.xml
RUN mvn dependency:go-offline -B
COPY ./src ./src
RUN mvn package && cp target/petstore-*.jar app.jar

FROM java:8-jre-alpine
MAINTAINER CyberArk

COPY --from=maven /app/app.jar /app.jar

ENTRYPOINT [ "java", "-jar", "/app.jar"]
