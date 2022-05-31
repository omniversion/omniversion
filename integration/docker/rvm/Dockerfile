FROM test-base

USER root
SHELL [ "/bin/bash", "-c" ]
RUN apt-add-repository -y ppa:rael-gc/rvm
RUN apt update && apt install rvm -y

SHELL [ "/bin/bash", "-i", "-l", "-c" ]
RUN rvm install 2.1.1
RUN rvm install 2.7.3
RUN rvm install 3.0.0
RUN rvm --default use 2.7.3
RUN rvm gemset create test1
RUN rvm 2.7.3@test1
RUN gem install nokogiri
RUN rvm gemset create test2
RUN rvm 2.7.3@test2
RUN gem install rails-api
