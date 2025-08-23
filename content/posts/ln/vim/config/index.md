+++
date = '2025-08-23T15:50:07+08:00'
draft = false
title = 'Vim config'
categories = ['技术']
tags=['VIM']
showHero = true
+++

Example:

{{< gallery >}}
    <img src="/learn/img/ln/vim/config/example.png" class="grid-w50"/>
{{< /gallery >}}

# 我的 VIM 配置

```
"============= base settings
" set search ignore the case
set ignorecase
" set file type match
syntax enable
syntax on

set nocompatible
filetype off

" set the help view
set nu
set cursorcolumn
set cursorline
set showmatch
set cc=120

" set the tab to space and default set 4 space
set tabstop=4
set softtabstop=4
set shiftwidth=4
set expandtab
set autoindent
set backspace=indent,eol,start

" vundle plugin manager
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

Plugin 'VundleVim/Vundle.vim'

" === the codes
" for golang
Plugin 'fatih/vim-go'


" === the managers
" for the file observe
Plugin 'preservim/nerdtree'
Plugin 'jistr/vim-nerdtree-tabs'
" Git status
Plugin 'Xuyuanp/nerdtree-git-plugin'

" === the helpers
" for bottom status
Plugin 'KeitaNakamura/neodark.vim'
Plugin 'crusoexia/vim-monokai'
Plugin 'acarapetis/vim-colors-github'
Plugin 'rakr/vim-one'
Plugin 'vim-airline/vim-airline'
Plugin 'vim-airline/vim-airline-themes'
Plugin 'jiangmiao/auto-pairs'
Plugin 'Valloric/YouCompleteMe', { 'on': [], 'commit': '85c11d3a875b02a7ac28fb96d0c7a02782f60410' }
Plugin 'airblade/vim-gitgutter'
Plugin 'vim-scripts/SuperTab'

" === for golangs
" params list, func list and so on, need https://github.com/jstemmer/gotags
Plugin 'majutsushi/tagbar'
" code jump
Plugin 'dgryski/vim-godef'

" === for dockerfile
Plugin 'ekalinin/Dockerfile.vim'

" === for mark down
Plugin 'iamcco/mathjax-support-for-mkdp'
Plugin 'iamcco/markdown-preview.vim'

call vundle#end()




filetype plugin indent on

" set auto nerdtree
autocmd VimEnter * NERDTree

" set key-map
" set the nerd tree find
nnoremap <C-f> :NERDTreeFind<CR>

" === plugin config
let g:gitgutter_terminal_reports_focus=0

"==============================================================================
" 主题配色
"==============================================================================

" 开启24bit的颜色，开启这个颜色会更漂亮一些
set termguicolors
" 配色方案, 可以从上面插件安装中的选择一个使用
colorscheme one " 主题
set background=dark " 主题背景 dark-深色; light-浅色


"==============================================================================
" vim-go 插件
"==============================================================================
let g:go_fmt_command = "goimports" " 格式化将默认的 gofmt 替换
let g:go_autodetect_gopath = 1
let g:go_list_type = "quickfix"

let g:go_version_warning = 1
let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_methods = 1
let g:go_highlight_generate_tags = 1

let g:godef_split=2
let g:godef_same_file_in_same_window=1

"==============================================================================
"  majutsushi/tagbar 插件
"==============================================================================

" majutsushi/tagbar 插件打开关闭快捷键
nnoremap Tt :TagbarToggle<CR>

let g:tagbar_type_go = {
    \ 'ctagstype' : 'go',
    \ 'kinds'     : [
        \ 'p:package',
        \ 'i:imports:1',
        \ 'c:constants',
        \ 'v:variables',
        \ 't:types',
        \ 'n:interfaces',
        \ 'w:fields',
        \ 'e:embedded',
        \ 'm:methods',
        \ 'r:constructor',
        \ 'f:functions'
    \ ],
    \ 'sro' : '.',
    \ 'kind2scope' : {
        \ 't' : 'ctype',
        \ 'n' : 'ntype'
    \ },
    \ 'scope2kind' : {
        \ 'ctype' : 't',
        \ 'ntype' : 'n'
    \ },
    \ 'ctagsbin'  : 'gotags',
    \ 'ctagsargs' : '-sort -silent'
\ }

"==============================================================================
"  nerdtree-git-plugin 插件
"==============================================================================
let g:NERDTreeGitStatusIndicatorMapCustom = {
    \ "Modified"  : "✹",
    \ "Staged"    : "✚",
    \ "Untracked" : "✭",
    \ "Renamed"   : "➜",
    \ "Unmerged"  : "═",
    \ "Deleted"   : "✖",
    \ "Dirty"     : "✗",
    \ "Clean"     : "✔︎",
    \ 'Ignored'   : '☒',
    \ "Unknown"   : "?"
    \ }

let g:NERDTreeGitStatusShowIgnored = 1



"==============================================================================
"  Valloric/YouCompleteMe 插件
"==============================================================================

" make YCM compatible with UltiSnips (using supertab)
let g:ycm_key_list_select_completion = ['<C-n>', '<space>']
let g:ycm_key_list_previous_completion = ['<C-p>', '<Up>']
let g:SuperTabDefaultCompletionType = '<C-n>'

" better key bindings for UltiSnipsExpandTrigger
let g:UltiSnipsExpandTrigger = "<tab>"
let g:UltiSnipsJumpForwardTrigger = "<tab>"
let g:UltiSnipsJumpBackwardTrigger = "<s-tab>"

let g:ycm_global_ycm_extra_conf = "~"
```
