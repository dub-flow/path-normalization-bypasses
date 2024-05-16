FROM node:alpine

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the package.json and package-lock.json files
COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 5000

# Define the command to run the app
CMD ["node", "app.js"]
