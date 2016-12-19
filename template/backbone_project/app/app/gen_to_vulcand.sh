#!/bin/bash

#echo $1 $2 $3 $4

USERNAME=jiajiachen

NAME=bgmodelset2

VULCAND=http://192.168.5.46:8182

LOCAL=http://192.168.1.19:8004

FRONTPATH='PathRegexp("'$USERNAME/$NAME'.*")'
#"'PathRegexp(\"/"$USERNAME"/"$NAME".*\")'"


vctl backend upsert -id b_$USERNAME_$NAME --vulcan $VULCAND

vctl server upsert -id svr_$USERNAME_$NAME -b b_$USERNAME_$NAME --url=$LOCAL --vulcan $VULCAND

echo $FRONTPATH

#echo "'PathRegexp(\"/${USERNAME}/${NAME}.*\")'"

vctl frontend upsert -id f_$USERNAME_$NAME --route $FRONTPATH -b b_$USERNAME_$NAME --vulcan $VULCAND

vctl rewrite upsert -f f_$USERNAME_$NAME -id r_$USERNAME_$NAME --regexp="/${USERNAME}/${NAME}/(.*)" --replacement='/$1' --vulcan $VULCAND
