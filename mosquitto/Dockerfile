FROM eclipse-mosquitto:latest

LABEL maintainer="KnT <khanhnd.it24@gmail.com>"

COPY config/mosquitto.conf /mosquitto/config/mosquitto.conf
COPY config/passwd /mosquitto/config/passwd

RUN mosquitto_passwd -U /mosquitto/config/passwd
