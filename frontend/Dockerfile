FROM node:24-alpine AS builder

WORKDIR /app

COPY . .

RUN npm install && npm run build

FROM nginx:stable-alpine as production-stage

# Copy the built files from the build-stage to the nginx html directory
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy custom nginx configuration files
COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY nginx/default.confg /etc/nginx/conf.d/default.conf

# Copy the entry point script
COPY start.sh /start.sh

# Make the entry point script executable
RUN chmod +x /start.sh

# Expose port 80
EXPOSE 80

# Set the entry point to the shell script
ENTRYPOINT ["/start.sh"]
