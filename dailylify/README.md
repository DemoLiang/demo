# dailyDemon
daily demon in daily time


how to create ssh key

1. ssh-keygen.exe -t rsa -C "[mail]" -b 4096 
2. ssh-agent -s
3. ssh-add ~/.ssh/id_rsa
4. add ssh public key to github
5. test your ssh key:ssh -T git@github.com

how to upload large file

1. you need to install lfs
2. git lfs install
3. git lfs track "file"
4. git add file
5. git commit
6. git push


how to create an review,issues,pull requests

1. git checkout -b [newbranch of you issues]
2. modify or fixed the bug
3. git add *
4. git commit -m "[fixed #4] this is the issues you want to fixed number relation"
5. git checkout master
6. git merge issues#4
7. git push
8. this step is to fixed the bug in master and relate the issue number

another step is create an pull request

1. git checkout -b [dev branch]
2. modify or fixed the bug
3. git add *
4. git commit -m "[fixed #5] this is the issues you want to fixed number relation issues"
5. git push origin [dev branch]
6. you can find the issues has fixed by the commit and relation modify code
7. create and pull request
