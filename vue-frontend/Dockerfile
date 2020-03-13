FROM node:lts-alpine as builder

RUN mkdir -p /app
WORKDIR /app

RUN npm install -g http-server

COPY ./package*.json ./
RUN npm install \
    # copy modules because they are es5, which can be included only from index.html
    && mkdir -p static/ \
    && cp node_modules/conic-gradient/conic-gradient.js static/ \
    && cp node_modules/prefixfree/prefixfree.min.js static/

COPY . .
RUN npm run build

FROM nginx:stable-alpine as product

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx_config/default.conf /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]
