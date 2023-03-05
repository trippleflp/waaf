FROM nginx
RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY conf/nginx.conf /etc/nginx/templates/nginx.conf.template
COPY /lib /etc/nginx/njs

EXPOSE 80

ENV NGINX_PORT=80
# ENV GROUP_ID=some_random_uuid
ENV NGINX_ENVSUBST_OUTPUT_DIR=/etc/nginx
