FROM node:15.11.0-alpine

# Define workdir
WORKDIR /app

# Copy build
COPY rollup.config.js ./
COPY package*.json ./

# Install dependancies
RUN npm install

# Copy source file
COPY ./src/    ./src/
COPY ./public/ ./public/

# Build svelte app
RUN npm run build

# Expose port 5000
EXPOSE 5000

# Run application on 0.0.0.0
ENV HOST=0.0.0.0

# Entry
ENTRYPOINT [ "npm", "run" ]
CMD [ "start" ]