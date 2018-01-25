# rackspaceKube

The terraform directory has the provisioning tf files to build out a simple 3 node cluster on the Google Cloud.

The kube directory has the kubernetes manifest files to load my image in the above cluster.

The key_value_thingy directory has the the golang code for this really basic microservice, 
it also has the dockerfile for making the image.

# Caveat
A true micro service would be scalable, and in this case, stateless.  The storage of the key value 
pairs would occur via api calls to a data tier application that supports concurrancy.  Database
technologies such as Redis or Mongdb would be perfect, but also could use any native technology
to the cloud environment, such as google cloud storage.  Any time in the code where we manipulate the
"Database" map, imagine instead an api call to remotely perform thos data operations.
