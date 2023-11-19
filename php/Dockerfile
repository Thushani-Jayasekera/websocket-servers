FROM php:cli

# Install system dependencies
RUN apt-get update && apt-get install -y \
    git \
    unzip \
    zip

# Install Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Allow Composer to run as super user
ENV COMPOSER_ALLOW_SUPERUSER=1

# Set up the working directory
WORKDIR /app

# Copy the application files
COPY . /app

# Install dependencies with Composer
RUN composer install

# Expose port 8009
EXPOSE 8009

# Start the WebSocket server
CMD ["php", "websocket_server.php"]
