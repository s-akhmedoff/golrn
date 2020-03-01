call plug#begin('~/.vim/plugged')

Plug 'fatih/vim-go' 

Plug 'scrooloose/syntastic'

Plug 'ajmwagar/vim-deus'

Plug 'tpope/vim-surround'

Plug 'vim-airline/vim-airline'

Plug 'honza/vim-snippets'

Plug 'vim-airline/vim-airline-themes'

call plug#end()

let g:airline#extensions#tabline#enabled = 1
let g:airline#extensions#tabline#formatter = 'unique_tail'

let g:airline_theme='deus'

set laststatus=2

syntax enable

set t_Co=256
set termguicolors

let &t_8f = "\<Esc>[38;2;%lu;%lu;%lum"
let &t_8b = "\<Esc>[48;2;%lu;%lu;%lum"

set background=dark    " Setting dark mode
colorscheme deus
let g:deus_termcolors=256

set noerrorbells
set novisualbell

set number

set noshowmode

filetype indent on
