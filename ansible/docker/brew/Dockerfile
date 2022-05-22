FROM test-base

USER root
RUN useradd -m -s /bin/bash linuxbrew && \
    echo 'linuxbrew ALL=(ALL) NOPASSWD:ALL' >>/etc/sudoers

USER linuxbrew
RUN curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh | bash

USER root
RUN chown -R ansible $(/home/linuxbrew/.linuxbrew/bin/brew --prefix)/*

USER ansible
ENV PATH="/home/linuxbrew/.linuxbrew/bin:${PATH}"
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"
RUN brew install git
RUN brew install node
