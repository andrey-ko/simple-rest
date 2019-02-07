param(
    [Parameter()]
    [ValidateSet("centos","win1809","win1803","win1709")]
    [string] $target = "centos"
)

$ErrorActionPreference = "Stop"

push-location $psScriptRoot

$dockerFile = gci "Dockerfile.$target"

docker build -t "akolomentsev/simple-rest:$target" -f $dockerFile .
docker push akolomentsev/simple-rest:$target

pop-location

