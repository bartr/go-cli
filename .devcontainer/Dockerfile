# this image is built and updated weekly
# https://github.com/cse-labs/codespaces-images

FROM ghcr.io/cse-labs/k3d:latest

# some images require specific values
ARG USERNAME=vscode

ENV PATH $PATH:/usr/local/bin/go:/workspaces/go-cli/go/bin
ENV GOPATH /workspaces/go-cli/go

RUN echo "$(date)    docker build started" >> /home/$USERNAME/status && \
    mkdir -p /home/vscode/go && \
    wget https://go.dev/dl/go1.17.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz && \
    rm -f go1.17.6.linux-amd64.tar.gz && \
    echo "compinit" >> /home/${USERNAME}/.zshrc && \
    echo "$(date)    docker build complete" >> /home/$USERNAME/status && \
    chown -R $USERNAME:$USERNAME /home/$USERNAME
