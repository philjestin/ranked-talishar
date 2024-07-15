
# Use the official node image with slim variant
FROM node:18-slim

# Set the working directory
WORKDIR /app

# Copy the package.json file
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Expose the port
EXPOSE 3000

RUN npm run build
ENV NODE_ENV=production
CMD ["npm", "run" ,"start"]