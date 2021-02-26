

By: He TianGang

# Repo Link

This repo can be found at the link below:

https://github.com/allenhe77/Calipsa_Task

# Metric Name

the name I assigned to my custom metric is:

"total"

# Issue with Chrome Browser

My chrome browser was sending 2 requests everytime I refreshed the page for go server, 
turns out it could be one of my chrome extensions is causing this, not sure if this
will be happening on your chrome browser as well.


# Namespace 

Both the golang pod and prometheus pod is confugured to run under namespace "monitoring"

"kubectl create namespace monitoring"


# Docker image url

https://hub.docker.com/repository/docker/effy77/calipsa_task

