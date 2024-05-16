FROM ruby:3.1-alpine

# Install build tools and libraries needed to build native extensions
RUN apk add --no-cache build-base

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the Gemfile and Gemfile.lock into the working directory
COPY Gemfile Gemfile.lock ./

# Install the gems defined in the Gemfile
RUN bundle install

# Copy the current directory contents into the container
COPY . .

EXPOSE 5000

# Run app.rb when the container launches
CMD ["ruby", "app.rb"]
