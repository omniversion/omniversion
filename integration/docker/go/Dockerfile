FROM test-base

USER root
RUN curl -s https://storage.googleapis.com/golang/go1.17.linux-amd64.tar.gz | sudo tar -v -C /usr/local -xz
RUN mkdir -p /repos/go
RUN chmod 777 /repos/go

USER ansible
ENV PATH $PATH:/usr/local/go/bin

# download some modules
WORKDIR /repos/go
RUN go get github.com/sirupsen/logrus
RUN go get github.com/juju/errors
RUN go install github.com/nathan-fiscaletti/example-go-install@latest
RUN git clone https://github.com/spf13/viper.git
WORKDIR /repos/go/viper
RUN go get
