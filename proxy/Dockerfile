FROM nginx:alpine

WORKDIR /proxy
COPY /templates/nginx.conf.template .

ARG PROXY_PORT
ARG FRONTEND_PORT
ARG BACKEND_PORT

RUN envsubst < /proxy/nginx.conf.template > /etc/nginx/conf.d/default.conf

ENTRYPOINT ["nginx", "-g", "daemon off;"]
