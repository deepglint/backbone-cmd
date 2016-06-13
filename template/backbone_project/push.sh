#!/usr/bin

#To build the project & to choose to tag it

tag=$1

echo "Doing the Tagging Release with $tag"

echo "Gulping the Project"

gulp

sleep 2s

gulp md

sleep 2s

echo "Pusing to Github"

git add .

git commit -m "add tag $tag"

git push

sleep 2s

echo "Tagging with tag name:  $tag"

git tag $tag

git push --tags

sleep 2s

echo "Done"

