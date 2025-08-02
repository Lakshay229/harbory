FROM node:18-alpine as build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

ARG VITE_API_URL
ARG VITE_APP_ENV
ARG VITE_APP_TITLE
ARG VITE_DEBUG_MODE
ARG VITE_API_TIMEOUT

ENV VITE_API_URL=$VITE_API_URL
ENV VITE_APP_ENV=$VITE_APP_ENV
ENV VITE_APP_TITLE=$VITE_APP_TITLE
ENV VITE_DEBUG_MODE=$VITE_DEBUG_MODE
ENV VITE_API_TIMEOUT=$VITE_API_TIMEOUT
RUN npm run build

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]
