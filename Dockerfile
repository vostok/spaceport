# build runtime image
FROM nginx:alpine

WORKDIR /usr/share/nginx/html
COPY --from=build-env /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
