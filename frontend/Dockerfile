FROM nginx:alpine

# Copy the custom Nginx config
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy the frontend files
COPY . /usr/share/nginx/html/

EXPOSE 80
