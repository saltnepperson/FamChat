# Build stage
FROM node:22-alpine AS build

WORKDIR /app

# Install dependencies
COPY ./famchat-ui/package*.json ./
RUN npm install -g pnpm

# Copy all source code and build the application
COPY ./famchat-ui .
RUN pnpm build


# Development stage
FROM node:22-alpine AS development

WORKDIR /app

# Copy everything from the build stage
COPY --from=build /app .

# Expose port 8000 for development
EXPOSE 8000

# Start the development server
CMD ["pnpm", "dev", "--port", "8000"]


# Production stage
FROM nginx:alpine AS production

WORKDIR /usr/share/nginx/html

# Copy the built files from the build stage to the Nginx html directory
COPY --from=build /app/.next .next
COPY --from=build /app/public public
COPY --from=build /app/package.json .

# Set environment variables for Next.js
ENV NODE_ENV=production

# Expose port 80 for Nginx
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
