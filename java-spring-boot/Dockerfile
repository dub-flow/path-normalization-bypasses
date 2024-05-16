# Use the official Maven image to create a build artifact.
# This uses Maven and OpenJDK 17, adjusted to match your project requirements
FROM maven:3.8.4-openjdk-17-slim AS build

# Copy the pom.xml file and download all dependencies.
COPY pom.xml /home/app/pom.xml
RUN mvn -f /home/app/pom.xml dependency:resolve

# Copy the project source
COPY src /home/app/src

# Build the application
RUN mvn -f /home/app/pom.xml clean package

# Use the official OpenJDK image for a base image.
FROM openjdk:17.0.1-jdk-slim

# Copy the jar file from the build stage into the /opt/app directory
COPY --from=build /home/app/target/some-test-0.0.1-SNAPSHOT.jar /opt/app/app.jar

# Set the working directory to /opt/app
WORKDIR /opt/app

# Expose the port the app runs on
EXPOSE 5000

# Set the container to execute the application on start
ENTRYPOINT ["java","-jar","app.jar"]
