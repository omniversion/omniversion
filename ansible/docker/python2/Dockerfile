FROM test-base

USER root
SHELL [ "/bin/bash", "-c"]
RUN add-apt-repository -y ppa:deadsnakes/ppa
RUN apt install -y python python-pip python-openssl
RUN pip install virtualenv

RUN mkdir -p /pip/test1
WORKDIR /pip/test1
RUN python -m virtualenv test-env
RUN bash -i -l -c "source test-env/bin/activate && pip install botocore && pip install PyYAML"

RUN mkdir -p /pip/test2
WORKDIR /pip/test2
RUN python -m virtualenv test-env
RUN bash -i -l -c "source test-env/bin/activate && pip install six"
