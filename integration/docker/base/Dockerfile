FROM ubuntu:bionic-20220427

LABEL maintainer="Layer9 GmbH <hello@layer9.berlin>"

RUN mkdir /repos
WORKDIR /repos


# USERS
# create an ansible user with passwordless sudo

# install sudo
RUN apt update && apt install sudo

# create the ansible user
ENV ANSIBLE_USER=ansible SUDO_GROUP=sudo DEPLOY_GROUP=deployer
RUN set -xe \
  && groupadd -r ${ANSIBLE_USER} \
  && groupadd -r ${DEPLOY_GROUP} \
  && useradd -ms "/bin/bash" -g ${ANSIBLE_USER} ${ANSIBLE_USER} \
  && usermod -aG ${SUDO_GROUP} ${ANSIBLE_USER} \
  && usermod -aG ${DEPLOY_GROUP} ${ANSIBLE_USER} \
  && sed -i "/^%${SUDO_GROUP}/s/ALL\$/NOPASSWD:ALL/g" /etc/sudoers


# APT
# install some useful packages
# both for the benefit of the derived containers and to test apt package fetching

# need to test pinned versions
COPY apt_preferences /etc/apt/preferences.d/apt_preferences

# install packages
ENV DEBIAN_FRONTEND=noninteractive
RUN apt update \
    &&  apt install -y --no-install-recommends \
        ansible \
        build-essential \
        ca-certificates \
        curl \
        file  \
        git \
        gnupg2 \
        libbz2-dev \
        libffi-dev \
        liblzma-dev \
        libncurses5-dev \
        libncursesw5-dev \
        libreadline-dev \
        libssl-dev \
        libsqlite3-dev \
        llvm \
        make \
        mysql-common \
        nginx \
        nmap \
        openssh-client \
        software-properties-common \
        tk-dev \
        wget \
        xz-utils \
        zlib1g-dev
