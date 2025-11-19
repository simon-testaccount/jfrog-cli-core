git status
git checkout -b test-branch
echo "Hello, world" > "test.txt"
git config --global user.email "you@example.com"
git config --global user.name "Your Name"
git add test.txt
git commit -m "Test commit"
git push -u origin test-branch
gh release edit v2.60.0 --notes-file release.md