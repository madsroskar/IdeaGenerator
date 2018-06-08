FROM golang

ARG app_env
COPY ./src /go/src/github.com/madsroskar/IdeaGenerator/src
WORKDIR /go/src/github.com/madsroskar/IdeaGenerator/src

ENV APP_ENV $app_env
ENV CWD /go/src/github.com/madsroskar/IdeaGenerator/src
ENV SOURCES_FILE $CWD/data/sources.txt
ENV TARGETS_FILE $CWD/data/targets.txt
ENV TMPL_DIR $CWD/tmpl/
ENV SRC_DIR $CWD

RUN echo $CWD


RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8000