git remote add origin git@bitbucket.org:Exotel/sif.git


# below is used to a add a new remote:
- git remote add origin git@github.com:User/UserRepo.git


# below is used to change the url of an existing remote repository:
- git remote set-url origin git@github.com:User/UserRepo.git


# below will push your code to the master branch of the remote repository defined with origin and -u let you point your current local branch to the remote master branch:
- git push -u origin master
- git push origin <localBranchName>:<RemoteBranchName>




