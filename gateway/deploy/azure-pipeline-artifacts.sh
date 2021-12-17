#!/bin/sh

value=`cat $1`
echo "$value" > $3
sed -i "s/latest/$2/g" $3