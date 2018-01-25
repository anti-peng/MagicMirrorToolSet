#!/bin/bash

PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/share/dotnet:/usr/local/go/bin
export PATH

# read -p "Please input your filename: " input
# subfix=$(node -v)
# filename=${input:-"defaultname"}${subfix}
# echo -e "${filename}"
# exit 0

# # chmod a+x magicmirror.sh; ./magicmirror.sh

# PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/share/dotnet:/usr/local/go/bin
# export PATH

# # start MagicMirror server
cd /Users/fry/Downloads/tmp/MagicMirror
node serveronly > ./server.log 2>&1
echo "done"
exit 0