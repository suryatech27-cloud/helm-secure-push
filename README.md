# Helm-Secure-Push
This Repository responsible for push chart into harbor repository.

Install this plugin : 
helm plugin install https://github.com/suryatech27-cloud/helm-secure-push

Support for Windows/amd64 --------> 
Support for linux/amd64 ---------->

secure-push -
flaglist :
1)username (-u) - username of a repository
2)password (-p) - Password of a repository
3)reponame (-repo) Repository name that you want to store chart
4)chartfile (-chart) Specify the chart file as a compressed with tar (.tgz)
5)provfile (-prov) Specify the provenance file of that chart (.tgz.prov) 
