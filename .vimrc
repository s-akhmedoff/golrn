call plug#begin('~/.vim/plugged')

Plug 'ajmwagar/vim-deus'

Plug 'preservim/nerdtree'

Plug 'vim-airline/vim-airline-themes'

Plug 'rafi/awesome-vim-colorschemes'

Plug 'tpope/vim-surround'

Plug 'vim-scripts/c.vim'

Plug 'airblade/vim-gitgutter'

Plug 'tpope/vim-fugitive'

Plug 'vim-airline/vim-airline'

Plug 'ryanoasis/vim-devicons'

Plug 'tpope/vim-surround'

Plug 'honza/vim-snippets'

Plug 'fatih/vim-go'

call plug#end()

let g:airline#extensions#tabline#enabled = 1 
let g:airline#extensions#tabline#formatter = 'unique_tail'

set t_Co=256
set termguicolors

let &t_8f = "\<Esc>[38;2;%lu;%lu;%lum"
let &t_8b = "\<Esc>[48;2;%lu;%lu;%lum"
 
set background=dark    " Setting dark mode
colorscheme deus
let g:deus_termcolors=256

filetype indent on
set noswapfile
set nobackup
set noerrorbells
syntax on
set laststatus=2
set number
set noshowmode
set encoding=UTF-8
map <C-m> :NERDTreeToggle<CR>
if (has("termguicolors"))
	set termguicolors
endif
colorscheme deus
let g:airline_powerline_fonts = 1
let g:airline_theme='deus'
