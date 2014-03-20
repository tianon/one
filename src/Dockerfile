FROM tianon/golang

# from https://www.kernel.org/
ENV KERNEL_VERSION 3.13.6

RUN curl -L https://www.kernel.org/pub/linux/kernel/v3.x/linux-$KERNEL_VERSION.tar.xz | tar -xJC /usr/src
RUN ln -sv linux-$KERNEL_VERSION /usr/src/linux

RUN apt-get update && apt-get install -yq bc

RUN ln -sfv ../one/src/kernel-config /usr/src/linux/.config

ADD . /usr/src/one/src
WORKDIR /usr/src/one/src
CMD ["./make.sh"]
