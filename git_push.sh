git config user.name "hgnulb"
git config user.email "hgnulb@163.com"
git config --local core.autocrlf false
git config --local core.quotepath false
git config --local advice.skippedCherryPicks false
git config --local gui.encoding utf-8
git config --local gui.logOutputEncoding utf-8
git config --local i18n.commitEncoding utf-8
git config --local alias.up "!git remote update -p; git merge --ff-only @{u}"

git fetch --all
git branch --set-upstream-to=origin/main main
git up
git add -A
git commit -m "no message" --allow-empty
git push -uf origin main