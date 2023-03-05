FROM nginx
RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY conf /etc/nginx
COPY lib /etc/njs

EXPOSE 8000
EXPOSE 443

CMD ["nginx","-g", "daemon off;"]