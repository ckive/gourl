# Use an official Node.js runtime as the base image
FROM node:16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the rest of the application files to the container
COPY . .

# Build the React app for production
RUN npm run build

# Install 'serve' globally
RUN npm install -g serve

# Expose port 3000 (default 'serve' port)
EXPOSE 3000

# Start 'serve' to serve the built React app
CMD ["serve", "-s", "build"]