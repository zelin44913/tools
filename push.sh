#!/bin/bash 

echo
echo "-------------------- Tools 仓库"
## 执行上传
git pull
git add -A .
git commit -m 'auto'
ver=`head -n 1 CHANGELOG.md | awk '{print $NF}'`
git tag -f "v${ver}"
git push --tags --force
git push

