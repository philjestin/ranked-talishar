# Use the official nginx image
FROM nginx:latest

# Copy the Nginx configuration file
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose the port
EXPOSE 80

# Run Nginx
CMD [ "nginx", "-g", "daemon off;" ]