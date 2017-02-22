# Shiba
This is the viewer of Github's contribution graph.

In Japan, that graph is called "Shiba".

# Screenshot
![Screenshot](https://raw.githubusercontent.com/0gajun/shiba/master/doc/screenshot.gif)

# Usage
```
# shiba [GitHub's user name] [--tz <timezone>]
```

# Configuration with environment variables
I recommend you that you set these environment variables to use shiba.

If you set environment variables, `SHIBA_GITHUB_USER_NAME` and `SHIBA_TIME_ZONE`,
you can show your shiba without any arguments.

* `SHIBA_GITHUB_USER_NAME`

  This value is default GitHub's user name.
  If you don't type any argument, shiba will show this user's contribution graph.

* `SHIBA_TIME_ZONE`

  This value specifies local time zone.
  Default value is `Asia/Tokyo`.
  If you set this value, shiba will automatically use it.

# Installation
```
# go get github.com/0gajun/shiba
```

# Author
0gajun <oga.ivc.s27@gmail.com>
