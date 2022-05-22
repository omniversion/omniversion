FROM test-base


USER root
SHELL ["/bin/bash", "-c"]
RUN git clone https://github.com/sstephenson/rbenv.git /root/.rbenv
RUN git clone https://github.com/sstephenson/ruby-build.git /root/.rbenv/plugins/ruby-build
RUN /root/.rbenv/plugins/ruby-build/install.sh
ENV PATH /root/.rbenv/bin:$PATH
RUN echo 'eval "$(rbenv init -)"' >> /etc/profile.d/rbenv.sh
RUN echo 'eval "$(rbenv init -)"' >> .bashrc
RUN touch /home/ansible/.bash_profile
RUN echo 'eval "$(rbenv init -)"' >> /home/ansible/.bash_profile
RUN source /home/ansible/.bash_profile
ENV CONFIGURE_OPTS --disable-install-doc

RUN rbenv install 2.7.1
RUN rbenv install 3.0.0
RUN rbenv global 2.7.1
SHELL ["/bin/bash", "-i", "-l", "-c"]
RUN gem update --system
RUN gem update
RUN gem install nokogiri -- --use-system-libraries
RUN gem install bundler rails-api
