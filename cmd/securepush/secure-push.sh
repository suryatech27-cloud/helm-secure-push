#!/bin/bash

#Taking user input

while [[ $# > 0 ]]
do
       key="$1"
               case $key in
                       --username )
                       uname="$2"
                       shift ;;
                       --password )
                       psword="$2"
                       shift ;;
                       --chart-tgz-file )
                       chartfile="$2"
                       shift ;;
                       --prov-file )
                       provfile="$2"
                       shift ;;
                       --repo-name )
                       repo="$2"
                       ;;

               esac
       shift
done

#Encoding the username with Base64 for Authorization
authtoken=$(echo -n $uname:$psword | base64)

#echo "username :" $uname
#echo "password : " $psword


#echo "Your packaged Signed  chart is :" $chartfile

#echo "Your Provenance file is :" $provfile

httpresponse=$(curl -s -w '%{response_code}' -H "authorization: Basic $authtoken" -H "Content-Type: multipart/form-data" -F "chart=@$chartfile; type=application/x-compressed-tar" -F "prov=@$provfile" -X POST "https://harbor.dell.com/api/chartrepo/$repo/charts")

#httpresponse=$(curl $verbose -s -u ":" -H "Content-Type: multipart/form-data" -F "chart=@$chartfile; type=application/x-compressed-tar" -F "prov=@$provfile" -X POST $repoUrl)

# get the length of the response
responseLength=`echo $httpresponse| wc -c`

if [ $responseLength -gt 3 ]; then
        httpStatusCode=${httpresponse:(-3)}
 else
        httpStatusCode=$httpresponse
fi

if [ $httpStatusCode -eq 201 ]; then

        echo "Done" 
else
        echo "Error Status :" $httpresponse
fi

#Finish

