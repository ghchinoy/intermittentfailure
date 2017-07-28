FROM scratch
MAINTAINER G. Hussain Chinoy <ghchinoy@gmail.com>
ADD intermittentfailure intermittentfailure
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/intermittentfailure"]