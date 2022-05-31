FROM test-base

USER ansible
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.2/install.sh | bash

SHELL ["/bin/bash", "-ilc"]
RUN nvm install 14.19.0
RUN nvm alias default 14.19.0
RUN nvm use default

RUN npm install -g pug

RUN mkdir -p /home/ansible/npm/test
WORKDIR /home/ansible/npm/test
RUN npm init --yes
RUN npm install async.js
