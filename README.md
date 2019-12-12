# diff-file
Show specified file's git diff

![](https://i.imgur.com/SQ8txFW.gif)

## Support OS
- Mac
- Linux

## Installtion
```sh
$ git clone https://github.com/skanehira/diff-file
$ cd
$ go install
```

## Usage
```sh
# specified file
$ diff-file gorilla.txt

# use with fuzzy finder
$ fzf | xargs diff-file
```

## Author
skanehira
