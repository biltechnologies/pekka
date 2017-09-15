# Pekka
Deploy and manage multiple wordpress sites with traefik and docker with Let's encrypt support. 

Pekka is a simple wrapper over `docker-compose`. Pekka generates and uses normal docker-compose files.  

# Prerequisites
docker and docker-compose should be installed and configured such that the user running `pekka` should have access and permissions to run `docker` and `docker-compose` commands

# Usage
1. Initialize pekka with the ``init`` command. 
  
   This creates the required files in the current directory.

    ```
    $ pekka init                     
    Enter traefik dashboard URL: <traefik dashboard URL>
    Enter let's encrypt email: <acme email>
    Creating network "pekkatraefik_webgateway" with driver "bridge"
    Creating pekkatraefik_proxy_1 ... 
    Creating pekkatraefik_proxy_1 ... done
    ```

    `<traefik dashboard URL>` is the URL where the traefik server's dashboard will be exposed.

    <acme email> is used for lets encrypt configurations.

2. Create a wordpress deployment using the ``create`` command

    ```
    $ pekka create toys
    Enter domain name: example.com
    Add entry for www.example.com? y
    Creating network "toys_default" with the default driver
    Pulling wordpress (nithinbose/wordpress:latest)...
    latest: Pulling from nithinbose/wordpress
    aa18ad1a0d33: Pull complete
    29d5f85af454: Pull complete
    eca642e7826b: Pull complete
    3638d91a9039: Pull complete
    3646a95ab677: Pull complete
    628b8373e193: Pull complete
    c24a2b2280ed: Pull complete
    f968b84cbbbc: Pull complete
    60fafe14064c: Pull complete
    bac57a95ddf1: Pull complete
    056ffd8ba0fc: Pull complete
    b595ac5a4e55: Pull complete
    5b72115923ec: Pull complete
    81b6cd799f34: Pull complete
    83faafba8a33: Pull complete
    577a4001244f: Pull complete
    69765c2499ed: Pull complete
    0044a72ca220: Pull complete
    5481d2b46462: Pull complete
    fcab5f51b65c: Pull complete
    0de0045cbc4b: Pull complete
    Digest: sha256:0f00bc21638db44478039e70e56ba40a0835b034a05300a4dcbfce2f86e26495
    Status: Downloaded newer image for nithinbose/wordpress:latest
    Creating toys_mysql_1 ... 
    Creating toys_mysql_1 ... done
    Creating toys_wordpress_1 ... 
    Creating toys_wordpress_1 ... done
    ```

    Pekka uses a custom wordpress docker image image ``nithinbose/wordpress``hosted on docker hub.

3. Remove a wordpress deployment using the ``remove`` command

    ```
    $ pekka remove toys
    Stopping toys_wordpress_1 ... done
    Stopping toys_mariadb_1 ... done
    Deployment stopped
    Going to remove toys_wordpress_1, toys_mariadb_1
    Removing toys_wordpress_1 ... done
    Removing toys_mariadb_1 ... done
    Removing files...
    Deployment removed
    ```