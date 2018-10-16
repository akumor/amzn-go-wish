FROM fedora:latest as downloader
WORKDIR /root
RUN dnf -y install git && \
    dnf clean all && \
    git clone https://github.com/doitlikejustin/amazon-wish-lister.git

RUN ls

FROM php:7.1.8-apache
MAINTAINER Jan Kumor

WORKDIR /srv/app
COPY --from=downloader /root/amazon-wish-lister/src /srv/app/
COPY vhost.conf /etc/apache2/sites-available/000-default.conf

RUN chown -R www-data:www-data /srv/app \
    && chmod -R 755 /srv/app \
    && a2enmod rewrite
