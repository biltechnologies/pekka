# Pekka
Deploy and manage multiple wordpress sites with traefik and docker with Let's encrypt support. 

Pekka is a simple wrapper over `docker-compose`. Pekka generates and uses normal docker-compose files.  

# Prerequisites
docker and docker-compose should be installed and configured such that the user running `pekka` should have access and permissions to run `docker` and `docker-compose` commands

# Usage
Initialize pekka with the ``init`` command. This creates the required files in the current directory.

```
$ ./pekka init                     
Enter traefik dashboard URL: <traefik dashboard URL>
Enter let's encrypt email: <acme email>
Creating network "pekkatraefik_webgateway" with driver "bridge"
Creating pekkatraefik_proxy_1 ... 
Creating pekkatraefik_proxy_1 ... done
```

`<traefik dashboard URL>` is the URL where the traefik server's dashboard will be exposed.

<acme email> is used for lets encrypt configurations.
