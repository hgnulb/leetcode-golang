git config user.name "hgnulb"
git config user.email "hgnulb@163.com"
git config core.autocrlf true
git checkout --orphan latest_branch
git add -A
git commit -m "no message"
git branch -D main
git branch -m main
git push -uf origin main
git branch --set-upstream-to=origin/main main
