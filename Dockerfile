FROM flyimg/flyimg-build
MAINTAINER mssz
COPY * /usr/local/bin/
RUN chmod +x /usr/local/bin/run.sh
RUN chmod +x /usr/local/bin/flyimg
CMD run.sh

