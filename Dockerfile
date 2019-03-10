FROM golang:1.11

RUN apt-get update
RUN apt-get install -y git tree vim curl

RUN curl https://raw.githubusercontent.com/skxeve/dotfiles/master/install.sh | bash
RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
RUN git clone https://github.com/fatih/vim-go.git ~/.vim/plugged/vim-go

RUN go get github.com/julienschmidt/httprouter

