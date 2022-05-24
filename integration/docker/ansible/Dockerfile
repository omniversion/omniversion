FROM test-base

USER root
SHELL [ "/bin/bash", "-c"]
RUN add-apt-repository -y ppa:deadsnakes/ppa
RUN apt install -y python3 python3-pip

USER ansible
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8
RUN python3 -m pip install -U pip setuptools
RUN python3 -m pip install --user -U ansible

RUN python3 -m ansible-galaxy install geerlingguy.nginx
RUN python3 -m ansible-galaxy collection install ansible.utils
