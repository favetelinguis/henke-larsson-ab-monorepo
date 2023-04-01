# Run the image in brew install --cask kitty

FROM ubuntu:22.10

# build-essential ,git, curl is needed for curl
RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    git

# Install brew
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
ENV PATH="/home/linuxbrew/.linuxbrew/bin:${PATH}"

# gcc is recomended by brew
RUN brew install \
    gcc \
    neovim \
    rtx \
    gh \
    lazygit \
    git-delta \
    ripgrep \
    fzf \
    zsh

# Setup zsh
RUN chsh -s /bin/zsh
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
ADD ./.p10k.zsh /root
ADD ./.zshrc /root

# Setup all config for tools work in my homerepo
WORKDIR /root
RUN mkdir -p /root/.config/{rtx,nvim,lazygit}

RUN git clone https://github.com/favetelinguis/dotfiles.git


RUN echo 'eval "$(rtx activate zsh)"' >> ~/.zshrc

ADD ./.rtx.toml /root
ADD ./.config.toml /root/.config/rtx

# Install all dev langs
RUN eval "$(rtx activate zsh)"
RUN rtx install

# run and setup zsh
