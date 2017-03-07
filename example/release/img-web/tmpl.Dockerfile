FROM 192.168.202.12:5000/xiaomei/nginx

COPY site.conf /etc/nginx/sites-enabled/{{ .ProName }}

WORKDIR /var/www/{{ .ProName }}
RUN mkdir /var/log/nginx/{{ .ProName }}

COPY public .
