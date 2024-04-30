[![Readme Card](https://github-readme-stats.vercel.app/api/pin/?username=cyclone-github&repo=passgen&theme=gruvbox)](https://github.com/cyclone-github/)

# Cyclone's PassGen
![image](https://i.imgur.com/9XzZVIm.png)

Password generator which creates passwords using upper / lower / digit & special char. Password length must be >= 8 char. Can be used to create 1 or more passwords by using CLI interface.

### Example Usage:
```
$ ./passgen.bin
 -----------------
| Cyclone PassGen |
 -----------------

Password length: 20
Number of passwords to generate: 10
,r/a/B6l.h0Ya#$\?dT.
c78i7/98"98*Y\G!]05<
fYex2u.a7\97dZ21829-
R"u51I8+]SP%]95*9o^^
9Y2'|47T=B1b.4\0hz53
f+Z89]J4v$l5}e2W<81I
<f2,1s0n3"o5V?(>?}]c
5z:7r1GSi(#f2[I=2g05
.O;&&at<1MC^76N.{{'4
*u15$4,A-5R9K]3OjX8p
```
### Compile from source:
- If you want the latest features, compiling from source is the best option since the release version may run several revisions behind the source code.
- This assumes you have Go and Git installed
  - `git clone https://github.com/cyclone-github/passgen.git`
  - `cd passgen`
  - `go mod init passgen`
  - `go mod tidy`
  - `go build -ldflags="-s -w" .`
- Compile from source code how-to:
  - https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt.txt
